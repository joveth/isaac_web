// other
package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/revel/revel"
)

type Other struct {
	App
}

func (c Other) Index() revel.Result {
	log.Debug("index")
	return c.RenderTemplate("App/Index.html")
}
