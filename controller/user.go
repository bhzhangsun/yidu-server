package controller

import "github.com/kataras/iris/v12"

func GetUserInfo(ctx iris.Context) {

	ctx.ViewData("UserInfo", "mobi")
	ctx.View("user.html")
}

func GetArticle(ctx iris.Context) {
	ctx.ViewData("Article", "llllll")
	ctx.View("article.html")
}
