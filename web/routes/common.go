package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.cool/web/controllers"
	"yidu.4se.cool/web/middleware"
)

func GetRouter(app *mvc.Application) {
	router := app.Router
	api := router.Party("/api")

	// 需要鉴权
	api.PartyFunc("/user", func(p iris.Party) {
		p.Post("/login", controllers.Login)
		p.Use(middleware.Auth)
		p.Get("/info", controllers.GetUserInfo)
		// p.Put("/info", controllers.ChangeInfo)
		p.Get("/logout", controllers.Logout)

	})
}
