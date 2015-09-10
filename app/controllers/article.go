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
	App
}

func (c Article) Index(page int) revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topics, pageno, totalPage := dao.GetTopicsWithTag(page, 3)
	page = pageno
	goods := dao.GetGoodTopics(0, 5)
	abouts := dao.GetGoodTopics(10, 5)
	return c.Render(topics, page, totalPage, goods, abouts)
}
func (c Article) Push() revel.Result {
	if _, ok := c.Session["user"]; !ok {
		c.Session["preUrl"] = "/push"
		c.Flash.Error("亲，登录之后才发布哦！")
		return c.Redirect("/login")
	}
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
	uName, ok := c.Session["user"]
	if !ok {
		c.Session["preUrl"] = "/push"
		c.Flash.Error("亲，登录之后才能发布哦！")
		return c.Redirect("/login")
	}
	topic.UName = uName
	dao, err := model.NewDao()
	topic.Status = 1
	if topic.Tag == 4 {
		topic.Status = 0
	}
	defer dao.Close()
	id, err := dao.InserTopic(topic)
	if err != nil {
		revel.ERROR.Printf("Unable to save topic: %v error %v", topic, err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	err = utils.UploadString(c.Request.PostFormValue("editormd-html-code"), fmt.Sprintf("%s.html", id))
	err = utils.UploadString(c.Request.PostFormValue("Body"), fmt.Sprintf("%s_code.html", id))
	return c.Redirect("/show/%s", id)
}

type CacheBean struct {
	HasCache bool
}

func (c Article) Show(id string) revel.Result {
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
	replays, repcnt := dao.GetReplays(id)
	return c.Render(topic, replays, repcnt)
}
func (c Article) Edit(id string) revel.Result {
	if _, ok := c.Session["user"]; !ok {
		c.Session["preUrl"] = fmt.Sprintf("/edit/%s", id)
		c.Flash.Error("亲，登录之后才能编辑哦！")
		return c.Redirect("/login")
	}
	dao, err := model.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	topic := dao.FindTopicById(id)
	if topic == nil {
		return c.Redirect("/")
	}
	content, ret := utils.GetCodeContent(id)
	if ret != nil {
		c.Response.Status = 500
		return c.RenderError(ret)
	}
	topic.Body = template.HTML(content)
	tags := dao.GetTags()
	return c.Render(topic, tags)
}
func (c Article) DoEdit(topic *model.Topic) revel.Result {
	//topic.Body = template.HTML(c.Request.PostFormValue("editormd-html-code"))
	_, ok := c.Session["user"]
	if !ok {
		c.Session["preUrl"] = fmt.Sprintf("/edit/%s", topic.Id)
		c.Flash.Error("亲，登录之后才能发布哦！")
		return c.Redirect("/login")
	}
	dao, err := model.NewDao()
	defer dao.Close()
	oldTopic := dao.FindTopicById(c.Request.PostFormValue("Id"))
	if oldTopic == nil {
		return c.Redirect("/")
	}
	oldTopic.Title = topic.Title
	oldTopic.Tag = topic.Tag
	if topic.Tag == 4 {
		oldTopic.Status = 0
	}
	id, err := dao.EditTopic(oldTopic)
	if err != nil {
		revel.ERROR.Printf("Unable to save topic: %v error %v", topic, err)
		c.Response.Status = 500
		return c.RenderError(err)
	}
	revel.INFO.Printf("The id: %s", id)
	utils.DeleteFile(id)
	err = utils.UploadString(c.Request.PostFormValue("editormd-html-code"), fmt.Sprintf("%s.html", oldTopic.Id.Hex()))
	err = utils.UploadString(c.Request.PostFormValue("Body"), fmt.Sprintf("%s_code.html", id))
	return c.Redirect("/show/%s", id)
}
func (c Article) DoComment(id string) revel.Result {
	user, ok := c.Session["user"]
	if !ok {
		c.Session["preUrl"] = fmt.Sprintf("/show/%s", id)
		c.Flash.Error("亲，登录之后才能回复哦！")
		return c.Redirect("/login")
	}
	replay := new(model.Replay)
	replay.TopicId = id
	replay.Content = c.Request.PostFormValue("content")
	replay.UName = user
	dao, _ := model.NewDao()
	defer dao.Close()
	topic := dao.FindTopicById(id)
	if topic == nil {
		return c.Redirect("/")
	}
	topic.Comment = topic.Comment + 1
	dao.InserReplay(replay)
	dao.EditTopic(topic)
	return c.Redirect("/show/%s", id)
}
