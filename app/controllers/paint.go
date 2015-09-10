package controllers

import (
	"github.com/revel/revel"
	"testapp/app/model"
)

type Paint struct {
	App
}

func (c Paint) Index(page int) revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		return c.Redirect("/")
	}
	topics, pageno, totalPage := dao.GetTopicsWithTag(page, 4)
	page = pageno
	goods := dao.GetGoodTopics(0, 5)
	abouts := dao.GetGoodTopics(10, 5)
	return c.Render(topics, page, totalPage, goods, abouts)
}
