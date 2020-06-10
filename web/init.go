package web

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.cool/web/routes"
)

var app *iris.Application

func Run() {
	app.Run(iris.Addr(":8000"), iris.WithOptimizations)
}

func init() {
	app = iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.Configure(app.Party("/app"), routes.RootRouteHandler)
}
