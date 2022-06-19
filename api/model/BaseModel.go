package model

import "time"

type BaseModel struct {
	createdTime time.Time `orm:"auto_now_add;type(datetime)"`
	updateTime  time.Time `orm:"auto_now;type(datetime)"`
	//updateTime  time.Time `orm:"auto_now;type(datetime);precision(4)"`
	del_flag int `orm:"default:"0"`
}
