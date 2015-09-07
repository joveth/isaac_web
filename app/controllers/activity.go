package controllers

import (
	"github.com/revel/revel"
)

type Activity struct {
	*revel.Controller
}

func (c Activity) Index() revel.Result {
	return c.Render()
}
