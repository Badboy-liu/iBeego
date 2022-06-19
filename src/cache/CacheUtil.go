package cache

import (
	"context"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/memcache"
	"time"
)

var (
	bm cache.Cache
)

func init() {
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

func Put(cache map[interface{}]interface{}) {
	for key, val := range cache {
		bm.Put(context.TODO(), key.(string), val, 10*time.Second)
	}
}
