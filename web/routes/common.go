package routes

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.cool/web/controllers"
)

func RootRouteHandler(app *mvc.Application) {
	router := app.Router
	mvc.Configure(router.Party("/api"), APIHandler)
}

func APIHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func AuthHandler(ctx iris.Context) {
	fmt.Println("Auth Operation")
	ctx.Next()
}
