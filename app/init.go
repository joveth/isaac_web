package app

import (
	"fmt"
	"github.com/revel/revel"
	"testapp/app/model"
	"time"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	//revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.TemplateFuncs["pls"] = func(a, b int) int { return a + b }
	revel.TemplateFuncs["mis"] = func(a, b int) int { return a - b }
	revel.TemplateFuncs["mo"] = func(a, b int) bool { return a%b == 0 }
	revel.TemplateFuncs["gt"] = func(a, b int) bool { return a > b }
	revel.TemplateFuncs["eq"] = func(a, b string) bool { return a == b }
	revel.TemplateFuncs["gTag"] = GetTag
	revel.TemplateFuncs["gUserLogo"] = GetUserLogo
	revel.TemplateFuncs["gUser"] = GetUserByName
	revel.TemplateFuncs["timesince"] = Timesince
	//revel.InterceptMethod((*controllers.App).CheckUser, revel.BEFORE)
}

func GetTag(id int) string {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		return ""
	}
	tag := dao.GetTag(id)
	return tag.Name
}
func GetUserLogo(name string) string {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		return ""
	}
	user := dao.GetUserLogoByName(name)
	if user.Logo == "" {
		return "default.jpg"
	}
	return user.Logo
}

func GetUserByName(name string) *model.User {
	dao, err := model.NewDao()
	defer dao.Close()
	if err != nil {
		return nil
	}
	user := dao.GetUserByName(name)
	return user
}
func Timesince(t time.Time) string {
	seconds := int(time.Since(t).Seconds())
	switch {
	case seconds < 60:
		return fmt.Sprintf("%d秒钟前", seconds)
	case seconds < 60*60:
		return fmt.Sprintf("%d分钟前", seconds/60)
	case seconds < 60*60*24:
		return fmt.Sprintf("%d小时前", seconds/(60*60))
	case seconds < 60*60*24*100:
		return fmt.Sprintf("%d天前", seconds/(60*60*24))
	default:
		return t.Format("2006-01-01")
	}
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func InitDB() {
	dao, err := model.NewDao()
	defer dao.Close()
	err = dao.InserTag(&model.Tag{1, "以撒のNews"})
	err = dao.InserTag(&model.Tag{2, "相关新闻"})
	err = dao.InserTag(&model.Tag{3, "同人文"})
	err = dao.InserTag(&model.Tag{4, "纯手绘"})
	err = dao.InserTag(&model.Tag{5, "漫画"})
	err = dao.InserTag(&model.Tag{6, "美图"})
	err = dao.InserTag(&model.Tag{7, "Mod"})
	err = dao.InserTag(&model.Tag{8, "Seed种子"})
	err = dao.InserTag(&model.Tag{9, "活动"})
	err = dao.InserTag(&model.Tag{10, "社区发展"})
	err = dao.InserTag(&model.Tag{11, "以撒的故事"})
	err = dao.InserTag(&model.Tag{12, "其他"})
	if err != nil {
	}
}
