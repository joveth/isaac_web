package controllers

import (
	"github.com/revel/revel"
)

type Paint struct {
	*revel.Controller
}

func (c Paint) Index() revel.Result {
	return c.Render()
}
