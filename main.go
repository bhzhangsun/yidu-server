package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"yidu.4se.tech/services/core/mfxsyd"
	"yidu.4se.tech/web/routes"
)

func main() {
	app := iris.New()

	novels := mfxsyd.NewNovels()
	novels.RegisterToClassifier()
	for i := 0; i < 2000; i++ {
		novels.Reducer(fmt.Sprintln("http://www.mianfeixiaoshuoyueduwang.com/book/%d", i))
	}

	//routes
	mvc.Configure(app.Party("/"), routes.RootRouteHandler)
	app.Run(iris.Addr(":8000"))
}
