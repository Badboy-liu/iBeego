// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"iBeego/controllers/test"
)

func RegisterController() {
	web.BConfig.RouterCaseSensitive = false
	web.BConfig.WebConfig.CommentRouterPath = "controllers/test"
	web.AutoRouter(&test.MainController{})
}
