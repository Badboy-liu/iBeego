package filter

import "github.com/beego/beego/v2/server/web"

type Advice interface {
	advice(filterFunc web.FilterFunc)
}
