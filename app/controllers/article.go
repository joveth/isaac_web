package controllers

import (
	"github.com/revel/revel"
)

type Article struct {
	*revel.Controller
}

func (c Article) Index() revel.Result {
	return c.Render()
}
