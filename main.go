package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/article", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.WriteString(path)

	})
	app.Run(iris.Addr("localhost:8000"))
}
