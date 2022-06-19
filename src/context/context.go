package context

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/beego/beego/v2/server/web/filter/opentracing"
	"github.com/beego/beego/v2/server/web/filter/prometheus"
	"github.com/beego/beego/v2/server/web/grace"
	_ "github.com/go-sql-driver/mysql"
	error2 "iBeego/controllers/error"
	IOrm "iBeego/orm"
	"iBeego/routers"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
	loadConfig()
	handlerCors()
	handlerRequest()
	initMysql()
	handlerError()
	IOrm.RegisterModel()
	routers.RegisterController()
	IOrm.CreateOrm()
}

func handlerError() {
	web.ErrorController(&error2.ErrorController{})
}

func handlerRequest() {
	web.BConfig.WebConfig.EnableXSRF = true
	web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	web.BConfig.WebConfig.XSRFExpire = 3600 //过期时间，默认1小时

	web.InsertFilterChain("/*", func(next web.FilterFunc) web.FilterFunc {
		return func(ctx *context.Context) {
			// do something
			logs.Info("hello")
			// don't forget this
			next(ctx)
			fb := &prometheus.FilterChainBuilder{}
			web.InsertFilterChain("/*", fb.FilterChain)

			opFb := &opentracing.FilterChainBuilder{}
			web.InsertFilterChain("/*", opFb.FilterChain)
			//tracer.SetGlobalTracer()
			// do something
		}
	})
}

func Run() {
	listen := web.BConfig.Listen
	web.Run(listen.HTTPSAddr, strconv.Itoa(listen.HTTPPort))
}

func loadConfig() {
	err := web.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatal("err:", err)
	}
}

func handlerCors() {
	web.InsertFilter("/*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	web.SetStaticPath("/static", "public")
}

func initMysql() {
	driverErr := orm.RegisterDriver("mysql", orm.DRMySQL)
	if driverErr != nil {
		log.Fatal("driverErr:", driverErr)
	}
	connectErr := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
	if connectErr != nil {
		log.Fatal("connectErr:", connectErr)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	err := grace.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
	w.Write([]byte("WORLD!"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
}
