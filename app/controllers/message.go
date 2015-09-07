package controllers

import (
	"github.com/revel/revel"
)

type Message struct {
	*revel.Controller
}

func (c Message) Index() revel.Result {
	return c.Render()
}
