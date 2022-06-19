package models

import (
	"reflect"
)

type User struct {
	Id   int64  `orm:"index"`
	Name string `orm:"column(name)"`
	//Name string `-`
	//Name string `orm:"size(60)"`
}

func (user User) TableName() string {
	return reflect.ValueOf(user).String()
}
