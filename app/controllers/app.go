package controllers

import (
	"github.com/revel/revel"
	"testapp/app/model"
)

type App struct {
	*revel.Controller
}

func (c App) Index(page int) revel.Result {
	revel.INFO.Printf("The page: %d", page)
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topics, pageno, totalPage := dao.GetTopics(page)
	page = pageno
	return c.Render(topics, page, totalPage)
}
func (c App) Mobile() revel.Result {
	return c.Render()
}
func (c App) Funcs() revel.Result {
	return c.RenderTemplate("App/Func.html")
}
func (c App) Login() revel.Result {
	return c.RenderTemplate("Login/Login.html")
}
func (c App) DoLogin() revel.Result {
	return c.Render()
}
func (c App) Register() revel.Result {
	return c.RenderTemplate("Login/Register.html")
}
func (c App) DoRegister() revel.Result {
	return  c.Redirect("/")
}