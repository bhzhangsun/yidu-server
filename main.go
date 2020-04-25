package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.tech/services/core/mfxsyd"
	"yidu.4se.tech/web/routes"
)

func main() {
	app := iris.New()

	novels := mfxsyd.NewNovels()
	novels.RegisterToClassifier()
	novels.Reducer("http://www.mianfeixiaoshuoyueduwang.com/book/1")

	//routes
	mvc.Configure(app.Party("/"), routes.RootRouteHandler)
	app.Run(iris.Addr(":8000"))
}
