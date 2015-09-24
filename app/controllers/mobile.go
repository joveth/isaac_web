package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"math/rand"
	"strconv"
	"testapp/app/model"
	"testapp/app/utils"
)

type Mobile struct {
	App
}

func (c Mobile) GetTopics() revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		return c.RenderJson("{}")
	}
	p := c.Request.PostFormValue("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}
	t := c.Request.PostFormValue("tag")
	tag, er := strconv.Atoi(t)
	if er != nil {
		tag = 1
	}
	revel.INFO.Printf("page: %s", p)
	topics := dao.GetTopicsWithTagForMob(page, tag)
	return c.RenderJson(topics)
}
func (c Mobile) GetUser() revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		return c.RenderJson("{}")
	}
	username := c.Request.PostFormValue("name")
	if username == "" {
		return c.RenderJson("{}")
	}
	revel.INFO.Printf("username: %s", username)
	user := dao.GetUserByName(username)
	return c.RenderJson(user)
}
func (c Mobile) Login() revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		return c.RenderJson("{}")
	}
	username := c.Request.PostFormValue("name")
	if username == "" {
		return c.RenderJson("{}")
	}
	pass := c.Request.PostFormValue("pass")
	if pass == "" {
		return c.RenderJson("{}")
	}
	revel.INFO.Printf("username: %s", username)
	user := dao.GetUserForLogin(username, pass)
	return c.RenderJson(user)
}
func (c Mobile) Regist() revel.Result {
	user := new(model.User)
	username := c.Request.PostFormValue("name")
	if username == "" {
		user.Email = "请输入用户名"
		return c.RenderJson(user)
	}
	email := c.Request.PostFormValue("email")
	if email == "" {
		user.Email = "请输入邮箱/手机"
		return c.RenderJson(user)
	}
	pass := c.Request.PostFormValue("pass")
	if pass == "" {
		user.Email = "请输入密码"
		return c.RenderJson(user)
	}
	user.Name = username
	user.Email = email
	user.Pass = pass
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		user.Name = ""
		user.Email = "数据保存失败了，请稍后再试！"
		return c.RenderJson(user)
	}
	hasUser := dao.GetUserByName(user.Name)
	if hasUser != nil && hasUser.Name != "" {
		user.Name = ""
		user.Email = "用户名已存在，请重新输入，或登录或找回密码！"
		return c.RenderJson(user)
	}
	hasUser = dao.GetUserByEmail(user.Email)
	if hasUser != nil && hasUser.Name != "" {
		user.Name = ""
		user.Email = "邮箱/手机 已存在，请重新输入，或登录或找回密码！"
		return c.RenderJson(user)
	}
	user.Flag = true
	user.Vip = 0
	user.Logo = fmt.Sprintf("default%d.jpg", rand.Intn(9))
	revel.ERROR.Printf("user: %v", user)
	err = dao.InserUser(user)
	if err != nil {
		user.Name = ""
		user.Email = "数据保存失败了，请稍后再试！"
		return c.RenderJson(user)
	}
	return c.RenderJson(user)
}
func (c Mobile) Show(id string) revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topic := dao.FindTopicById(id)
	if topic == nil {
		return c.Redirect("/")
	}
	if topic.Status == 0 {
		return c.RenderTemplate("Article/result.html")
	}
	content, ret := utils.GetHTMLContent(id)
	if ret != nil {
		c.Response.Status = 500
		return c.RenderError(ret)
	}
	revel.INFO.Printf("content: %v", content)
	return c.RenderHtml(content)
}
func (c Mobile) GetTwitter() revel.Result {
	content, _ := utils.GetTwitterHTML("https://twitter.com/MesutOzil1088")
	return c.RenderHtml(content)
}
func (c Mobile) GetFacebook() revel.Result {
	content, _ := utils.GetHTMLContentWithURL("https://www.facebook.com/mesutoezil?fref=nf")
	return c.RenderHtml(content)
}
