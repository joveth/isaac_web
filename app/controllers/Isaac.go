package controllers

import (
	"github.com/revel/revel"
)

type Isaac struct {
	App
}

func (c Isaac) Index() revel.Result {
	return c.Render()
}
