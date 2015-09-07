package controllers

import (
	"github.com/revel/revel"
)

type Isaac struct {
	*revel.Controller
}

func (c Isaac) Index() revel.Result {
	return c.Render()
}
