package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"html/template"
	"testapp/app/model"
	"testapp/app/utils"
)

type Article struct {
	*revel.Controller
}

func (c Article) Index() revel.Result {
	return c.Render()
}
func (c Article) Push() revel.Result {
	return c.Render()
}
func (c Article) DoPush(topic *model.Topic) revel.Result {
	//topic.Body = template.HTML(c.Request.PostFormValue("editormd-html-code"))
	topic.UName = "joveth"
	topic.Tag = "1"
	fmt.Printf("result=%v", topic.Body)
	dao, err := model.NewDao()
	defer dao.Close()
	id, err := dao.InserTopic(topic)
	if err != nil {
		fmt.Printf("err=%v", err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	err = utils.UploadString(c.Request.PostFormValue("editormd-html-code"), fmt.Sprintf("%s.html", id))
	return c.Redirect("/show/%s", id)
}

func (c Article) Show(id string) revel.Result {
	fmt.Printf("id=%v", id)
	dao, err := model.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topic := dao.FindTopicById(id)
	content, ret := utils.GetHTMLContent(id)
	if ret != nil {
		c.Response.Status = 500
		return c.RenderError(ret)
	}
	topic.Body = template.HTML(content)
	defer dao.Close()
	return c.Render(topic)
}
