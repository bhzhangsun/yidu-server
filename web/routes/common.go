package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mobi.4se.tech/web/controllers"
)

// var rootRouter router.Party

func RootRouteHandler(app *mvc.Application) {
	rootRouter := app.Router
	mvc.Configure(rootRouter.Party("/api"), APIHandler)
}

func APIHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func AuthHandler(ctx iris.Context) {
	fmt.Println("Auth Operation")
	ctx.Next()
}
