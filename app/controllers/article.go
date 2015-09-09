package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"html/template"
	"strings"
	"testapp/app/model"
	"testapp/app/utils"
	"time"
)

type Article struct {
	*revel.Controller
}

func (c Article) Index() revel.Result {
	return c.Render()
}
func (c Article) Push() revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	tags := dao.GetTags()
	return c.Render(tags)
}
func (c Article) DoPush(topic *model.Topic) revel.Result {
	//topic.Body = template.HTML(c.Request.PostFormValue("editormd-html-code"))
	topic.UName = "joveth"
	dao, err := model.NewDao()
	defer dao.Close()
	id, err := dao.InserTopic(topic)
	if err != nil {
		revel.ERROR.Printf("Unable to save topic: %v error %v", topic, err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	err = utils.UploadString(c.Request.PostFormValue("editormd-html-code"), fmt.Sprintf("%s.html", id))
	return c.Redirect("/show/%s", id)
}

type CacheBean struct {
	HasCache bool
}

func (c Article) Show(id string) revel.Result {
	dao, err := model.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topic := dao.FindTopicById(id)
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	revel.INFO.Printf("Ip : %s", ip)
	var obj CacheBean
	key := fmt.Sprintf("%s.%s", ip, id)
	err = cache.Get(key, &obj)
	if err == nil {
		revel.INFO.Printf("Cache : %v", obj)
	} else {
		revel.ERROR.Printf("Cache Errr : %v", err)
		if !obj.HasCache {
			obj.HasCache = true
			go cache.Set(key, obj, 10*time.Minute)
			topic.Read = topic.Read + 1
			dao.UpdateTopic(topic)
		}
	}
	content, ret := utils.GetHTMLContent(id)
	if ret != nil {
		c.Response.Status = 500
		return c.RenderError(ret)
	}

	topic.Body = template.HTML(content)
	defer dao.Close()
	return c.Render(topic)
}
