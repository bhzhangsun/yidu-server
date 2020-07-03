package web

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"yidu.4se.cool/web/routes"
)

var app *iris.Application

var cookie = securecookie.New(securecookie.GenerateRandomKey(16), securecookie.GenerateRandomKey(16))
var session = sessions.New(sessions.Config{
	Cookie:       "conect-token",
	Encode:       cookie.Encode,
	Decode:       cookie.Decode,
	AllowReclaim: true,
})

func Run() {
	app.Run(iris.Addr(":8000"), iris.WithOptimizations)
}

func init() {
	app = iris.New()
	app.Use(session.Handler())
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.Configure(app.Party("/"), routes.GetRouter)
}
