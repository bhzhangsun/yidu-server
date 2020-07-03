package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func Auth(ctx iris.Context) {
	sess := sessions.Get(ctx)
	if sess.Get("user") != nil {
		ctx.Next()
	}
}
