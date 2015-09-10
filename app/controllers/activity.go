package controllers

import (
	"github.com/revel/revel"
	"testapp/app/model"
)

type Activity struct {
	App
}

func (c Activity) Index(page int) revel.Result {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		revel.ERROR.Printf("Unable to open db:error %v", err)
		return c.Redirect("/")
	}
	topics, pageno, totalPage := dao.GetTopicsWithTag(page, 9)
	page = pageno
	goods := dao.GetGoodTopics(0, 5)
	return c.Render(topics, page, totalPage, goods)
}
