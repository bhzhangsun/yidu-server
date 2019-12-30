package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mobi.4se.tech/routes"
	"mobi.4se.tech/utils"
)

func main() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")

	conf, err := utils.GetConfiguration("./config/config.json")
	if err != nil {
		fmt.Println(err)
	}

	db, err := utils.GetDB((*conf).DB)
	if err != nil {
		fmt.Println(err)
	}

	db.Ping()

	tmpl.Reload(true)
	app.RegisterView(tmpl)

	//routes
	mvc.Configure(app.Party("/"), routes.RootRouteHandler)

	app.Run(iris.Addr("localhost:8000"))
}
