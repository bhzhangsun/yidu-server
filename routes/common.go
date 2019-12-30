package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mobi.4se.tech/controllers"
)

// var rootRouter router.Party

func RootRouteHandler(app *mvc.Application) {
	rootRouter := app.Router
	mvc.Configure(rootRouter.Party("/api"), APIHandler)
	mvc.Configure(rootRouter.Party("/admin"), AdminHandler)
	mvc.Configure(rootRouter.Party("/web"), WebHandler)
	mvc.Configure(rootRouter.Party("/wechat"), WechatHandler)
}

func APIHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func AdminHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func WebHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func WechatHandler(app *mvc.Application) {
	app.Router.Use(AuthHandler)
	mvc.New(app.Router.Party("/user")).Handle(new(controllers.UserController))
}

func AuthHandler(ctx iris.Context) {
	fmt.Println("Auth Operation")
	ctx.Next()
}

func handlerAPI(ctx iris.Context) {
	fmt.Println("handlerAPI")
	ctx.Next()
}

func handlerUser(ctx iris.Context) {
	fmt.Println("handlerUser")
	ctx.Next()
}
