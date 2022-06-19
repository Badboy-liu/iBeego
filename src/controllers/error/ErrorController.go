package error

import "github.com/beego/beego/v2/server/web"

type ErrorController struct {
	web.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
