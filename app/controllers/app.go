package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"math/rand"
	"testapp/app/model"
)

type App struct {
	*revel.Controller
}

func (c App) Index(page, tag int) revel.Result {
	revel.INFO.Printf("The page: %d", page)
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	goods := dao.GetGoodTopics(0, 5)
	abouts := dao.GetGoodTopics(10, 5)
	if tag > 0 {
		topics, pageno, totalPage := dao.GetTopicsWithTag(page, tag)
		page = pageno
		return c.Render(topics, page, totalPage, tag, goods, abouts)
	} else {
		topics, pageno, totalPage := dao.GetTopics(page)
		page = pageno
		return c.Render(topics, page, totalPage, tag, goods, abouts)
	}
}
func (c App) Mobile() revel.Result {
	return c.Render()
}
func (c App) Funcs() revel.Result {
	return c.RenderTemplate("App/Func.html")
}
func (c App) Login() revel.Result {
	if _, ok := c.Session["user"]; ok {
		return c.Redirect("/")
	}
	return c.RenderTemplate("Login/Login.html")
}
func (c App) DoLogin(user *model.User) revel.Result {
	if user.Name != "" && user.Pass != "" {
		dao, err := model.NewDao()
		defer dao.Close()
		if err != nil {
			c.Flash.Error("用户登录失败，请稍后重试！")
			return c.Redirect("/login")
		}
		loginUser := dao.GetUserForLogin(user.Name, user.Pass)
		if loginUser == nil || loginUser.Name == "" {
			c.Flash.Error("用户名或密码错误！")
			return c.Redirect("/login")
		}
		c.Session["user"] = loginUser.Name
		c.RenderArgs["theUser"] = loginUser
		if preUrl, ok := c.Session["preUrl"]; ok {
			delete(c.Session, "preUrl")
			return c.Redirect(preUrl)
		}
		return c.Redirect("/")
	}
	c.Flash.Error("用户名或密码错误！")
	return c.Redirect("/login")
}
func (c App) Register() revel.Result {
	if _, ok := c.Session["user"]; ok {
		return c.Redirect("/")
	}
	return c.RenderTemplate("Login/Register.html")
}
func (c App) DoRegister(user *model.User) revel.Result {
	user.Flag = false
	user.Vip = 0
	user.Logo = fmt.Sprintf("default%d.jpg", rand.Intn(9))
	//todo check
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		c.Flash.Error("用户保存失败，请稍后重试！")
		return c.Redirect("/register")
	}
	hasUser := dao.GetUserByName(user.Name)
	if hasUser != nil && hasUser.Name != "" {
		c.Flash.Error("用户名已存在，请登录或找回密码！")
		return c.Redirect("/register")
	}
	hasUser = dao.GetUserByEmail(user.Email)
	if hasUser != nil && hasUser.Name != "" {
		c.Flash.Error("邮箱/手机 已存在，请登录或找回密码！")
		return c.Redirect("/register")
	}
	err = dao.InserUser(user)
	if err != nil {
		c.Flash.Error("用户保存失败，请稍后重试！")
		return c.Redirect("/register")
	}
	c.Session["user"] = user.Name
	return c.Redirect("/")
}
func (c *App) CheckUser() revel.Result {
	user := c.user()
	if user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c *App) user() *model.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*model.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c *App) getUser(username string) *model.User {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		return nil
	}
	hasUser := dao.GetUserByName(username)
	if hasUser != nil && hasUser.Name != "" {
		return hasUser
	}
	return nil
}
func (c App) ShowUser(name string) revel.Result {
	if name == "" {
		return c.Redirect("/")
	}
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		return c.Redirect("/")
	}
	userShow := dao.GetUserByName(name)
	topics := dao.GetTopicsByUserName(name, 10)
	return c.Render(userShow, topics)
}
func (c App) ChoosePhoto(name string) revel.Result {
	if username, ok := c.Session["user"]; ok {
		if name == "" {
			return c.Redirect("/")
		}
		if username != name {
			return c.Redirect("/user?name=%s", name)
		}
		dao, err := model.NewDao()
		defer dao.Close()
		if err != nil {
			return c.Redirect("/")
		}
		user := dao.GetUserByName(name)
		if user == nil {
			return c.Redirect("/")
		}
		user.Logo = c.Request.PostFormValue("defaultAvatars")
		dao.UpdateUser(user)
		return c.Redirect("/user?name=%s", name)
	} else {
		c.Session["preUrl"] = fmt.Sprintf("/user?name=%s", name)
		c.Flash.Error("亲，登录超时，请重新登录！")
		return c.Redirect("/login")
	}
}
