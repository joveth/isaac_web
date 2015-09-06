package controllers

import (
	. "github.com/qiniu/api.v6/conf"
	qio "github.com/qiniu/api.v6/io"
	"github.com/qiniu/api.v6/rs"
	"github.com/revel/revel"
	"io"
	"testapp/app/db"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	dao,err := db.NewDao()
	defer dao.Close()
	collect :=dao.session.DB("isaac").C("user")
	err = collect.Insert(&model.User{"joveth", "joveth1@163.com","123456","123456"},
	               &model.User{"joveth1", "joveth1@163.com","123456","123456"})
        if err != nil {
                log.Fatal(err)
        }
	return c.RenderTemplate("App/Index.html")
}
func (c App) Mobile() revel.Result {
	return c.Render()
}
func (c App) Funcs() revel.Result {
	return c.RenderTemplate("App/Func.html")
}
func (c App) Upload(file io.Reader) revel.Result {
	var err error
	var ret qio.PutRet
	ACCESS_KEY = "NPfcHtb0e2EH7lCmmJot21MRr0lCel81S-QlUaJF"
	SECRET_KEY = "6DzF_oVRYhBkq0mqb4txThza_IfQEUey107VXaPq"

	var policy = rs.PutPolicy{
		Scope: "isaac",
	}

	err = qio.Put(nil, &ret, policy.Token(nil), "111", file, nil)

	if err != nil {
		revel.ERROR.Println("io.Put failed:", err)
		return c.RenderJson(map[string]string{
			"state": "UNKNOWN",
		})
	} else {
		return c.RenderJson(map[string]string{
			"title":    "",
			"original": "111",
			"state":    "SUCCESS",
		})
	}
}
