package controllers

import (
	"github.com/revel/revel"
)

type Message struct {
	App
}

func (c Message) Index() revel.Result {
	return c.Render()
}
