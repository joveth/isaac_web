package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//	dao, err := model.NewDao()
	//	defer dao.Close()
	//	err = dao.InserUser(&model.User{"joveth", "joveth1@163.com", "123456", "123456", time.Now(), "1"})
	//	if err != nil {
	//
	//	}
	return c.Render()
}
func (c App) Mobile() revel.Result {
	return c.Render()
}
func (c App) Funcs() revel.Result {
	return c.RenderTemplate("App/Func.html")
}
