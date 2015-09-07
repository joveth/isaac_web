// other
package controllers

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/revel/revel"
)

type Other struct {
	*revel.Controller
}

func (c Other) Index() revel.Result {
	log.Debug("index")
	fmt.Print("other.index")
	return c.RenderTemplate("App/Index.html")
}
