package base

import (
	"encoding/json"
	"encoding/xml"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-macaron/i18n"
	"gopkg.in/yaml.v3"
	"iBeego/models"
	"strconv"
	"strings"
	"time"
)

type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	web.Controller
	i18n.Locale
	user    models.User
	isLogin bool
}

func (this BaseController) ToJson(any interface{}) {
	data, _ := json.Marshal(any)
	this.setDate("json", 200, data)

}

func (this BaseController) ToXml(any interface{}) {
	data, _ := xml.Marshal(any)
	this.setDate("xml", 200, data)
}
func (this BaseController) ToJavaScript(any interface{}) {
	data, _ := json.Marshal(any)
	this.setDate("jsonp", 200, data)

}

func (this BaseController) ToYaml(any interface{}) {
	data, _ := yaml.Marshal(any)
	this.setDate("yaml", 200, data)

}

func (this BaseController) ToError(any interface{}) {
	data, _ := json.Marshal(any)
	this.setDate("json", 500, data)

}
func (this BaseController) setDate(format string, code int, data []byte) {
	this.Data[format] = map[string]string{
		"code": strconv.Itoa(code),
		"data": strings.ReplaceAll(string(data), "\"", ""),
	}
	switch format {
	case "json":
		this.ServeJSON(false)
	case "xml":
		this.ServeXML()
	case "yaml":
		this.ServeYAML()
	case "jsonp":
		this.ServeJSONP()
	}

}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	//this.Data["AppDescription"] = utils.AppDescription
	//this.Data["AppKeywords"] = utils.AppKeywords
	//this.Data["AppName"] = utils.AppName
	//this.Data["AppVer"] = utils.AppVer
	//this.Data["AppUrl"] = utils.AppUrl
	//this.Data["AppLogo"] = utils.AppLogo
	//this.Data["AvatarURL"] = utils.AvatarURL
	//this.Data["IsProMode"] = utils.IsProMode

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
