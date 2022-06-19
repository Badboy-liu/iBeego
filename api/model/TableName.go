package model

type Name interface {
	TableName(model interface{}) string
}
