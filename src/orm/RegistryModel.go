package orm

import (
	"github.com/beego/beego/v2/client/orm"
	"iBeego/models"
	"io"
)

var (
	MOrm orm.Ormer
)

func SetConfig() {
	//orm.RegisterModelWithPrefix("prefix_", new(interface{}))
	//orm.Debug=true
	//orm.ResetModelCache()

}

func logger() {
	// 默认名称  关闭表后再建表   打印执行过程
	orm.RunSyncdb("default", true, true)
	var w io.Writer
	orm.NewLog(w)
}

func RegisterModel() {
	orm.RegisterModel(&models.User{})
}

func CreateOrm() {
	MOrm = orm.NewOrm()
	MOrm.DBStats()
}
