package vo

type UserDo struct {
	Id   int         `from:"_"`
	Name interface{} `from:"name"`
}
