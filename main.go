package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.tech/web/routes"
)

func main() {
	app := iris.New()

	//routes
	mvc.Configure(app.Party("/"), routes.RootRouteHandler)
	app.Run(iris.Addr(":8000"))
}
