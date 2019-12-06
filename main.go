package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"mobi.4se.tech/controller"
	_ "mobi.4se.tech/service"
)

func main() {
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")

	tmpl.Reload(true)
	app.RegisterView(tmpl)

	api := app.Party("/api", handlerAPI)
	{
		user := api.Party("/user", handlerUser)
		user.Get("/info", controller.GetUserInfo)
		user.Get("/article", controller.GetArticle)
	}

	app.Get("/article", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.WriteString(path)

	})
	app.Run(iris.Addr("localhost:8000"))
}

func handlerAPI(ctx iris.Context) {
	fmt.Println("handlerAPI")
	ctx.Next()
}

func handlerUser(ctx iris.Context) {
	fmt.Println("handlerUser")
	ctx.Next()
}
