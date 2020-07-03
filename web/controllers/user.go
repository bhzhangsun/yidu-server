package controllers

import (
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"yidu.4se.cool/web/utils"
)

func Login(ctx iris.Context) {
	type Param struct {
		Phone string `json: "phone"`
		Code  string `json: "code"`
	}
	var param Param
	if ctx.ReadJSON(&param) != nil {
		ctx.JSON(&utils.Response{
			Code:    utils.RESPCODE_ERROR_PARAM,
			Data:    nil,
			Message: utils.RESPMSG_ERROR_PARAM,
		})
	} else {
		sess := sessions.Get(ctx)

		sess.Set("user", "")
		ctx.WriteString("dfdfrr")
	}
}

func GetUserInfo(ctx iris.Context) {
	sess := sessions.Get(ctx)

	ctx.JSON(iris.Map{
		"user": sess.Get("key"),
	})
}

func Logout(ctx iris.Context) {
	sess := sessions.Get(ctx)
	sess.Destroy()
}

func GetArticle(ctx iris.Context) {
	ctx.ViewData("Article", "llllll")
	ctx.View("article.html")
}

// UserController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type UserController struct{}

var helloView = mvc.View{
	Name: "user.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

// Get will return a predefined view with bind data.
//
// `mvc.Result` is just an interface with a `Dispatch` function.
// `mvc.Response` and `mvc.View` are the builtin result type dispatchers
// you can even create custom response dispatchers by
// implementing the `github.com/kataras/iris/hero#Result` interface.
func (c *UserController) Get() mvc.Result {

	// ok := services.NewUserService().Register(&models.User{
	// 	UserId: "haha",
	// 	Name:   "baohong",
	// })

	return mvc.View{
		Name: "user.html",
		Data: map[string]interface{}{
			"Title":     "defd",
			"MyMessage": "Welcome to my awesome website",
		},
	}
}

// you can define a standard error in order to re-use anywhere in your app.
var errBadName = errors.New("bad name")

// you can just return it as error or even better
// wrap this error with an mvc.Response to make it an mvc.Result compatible type.
var badName = mvc.Response{Err: errBadName, Code: 400}

// GetBy returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/hello/iris
// curl -i http://localhost:8080/hello/anything
func (c *UserController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
		// or
		// GetBy(name string) (mvc.Result, error) {
		//	return nil, errBadName
		// }
	}

	// return mvc.Response{Text: "Hello " + name} OR:
	return mvc.View{
		Name: "hello/name.html",
		Data: name,
	}
}
