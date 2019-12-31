package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mobi.4se.tech/web/routes"
)

func main() {
	app := iris.New()
	tmpl := iris.HTML("./web/views", ".html")

	tmpl.Reload(true)
	app.RegisterView(tmpl)

	//routes
	mvc.Configure(app.Party("/"), routes.RootRouteHandler)

	app.Run(iris.Addr("localhost:8000"))
}
