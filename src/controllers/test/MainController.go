package test

import (
	"api/vo"
	"encoding/json"
	"iBeego/controllers/base"
	"log"
	"mime/multipart"
)

type MainController struct {
	base.BaseController
}

//func (this *MainController) URLMapping() {
//	this.Mapping("Get", this.Get)
//}

//  @router /Main/GetInfo [get]
func (this *MainController) GetInfo() {
	userDo := vo.UserDo{}
	err := this.ParseForm(&userDo)
	if err != nil {
		return
	}
	this.ToJson(`你好`)
	this.ToError(`你好`)
}

// Add @router /test/:Add [post]
func (this *MainController) Add() {
	userDo := vo.UserDo{}
	err := this.ParseForm(&userDo)
	if err != nil {
		return
	}
	this.Ctx.WriteString("hello")
}

// Update @router /test/:Update [post]
func (this *MainController) Update() {
	userDo := vo.UserDo{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userDo)
	if err != nil {

	}
	if err != nil {
		return
	}
	this.Ctx.WriteString("hello")
}

// Upload @router /test/:Upload [post]
func (this *MainController) Upload() {
	file, head, err := this.GetFile("uploadName")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(file multipart.File) {
		fileErr := file.Close()
		if err != nil {
			log.Fatal("err:", fileErr)
		}
	}(file)
	// 保存位置在 static/upload, 没有文件夹要先创建
	fileErr := this.SaveToFile("uploadname", "static/upload/"+head.Filename)
	if err != nil {
		log.Fatal("err:", fileErr)
	}
}

// UrlParse @router /test/:UrlParse [get]
func (this *MainController) UrlParse() {
	// ?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie

	var id int
	this.Ctx.Input.Bind(&id, "id") //id ==123

	var isok bool
	this.Ctx.Input.Bind(&isok, "isok") //isok ==true

	var ft float64
	this.Ctx.Input.Bind(&ft, "ft") //ft ==1.2

	ol := make([]int, 0, 2)
	this.Ctx.Input.Bind(&ol, "ol") //ol ==[1 2]

	ul := make([]string, 0, 2)
	this.Ctx.Input.Bind(&ul, "ul") //ul ==[str array]

	//user =={Name:"astaxie"}
	user := struct{ Name string }{}
	err := this.Ctx.Input.Bind(&user, "user")
	if err != nil {
		log.Fatal("err:", err)
	}
}

// Delete @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *MainController) Delete() {
	//uid := u.GetString(":uid")
	//models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
