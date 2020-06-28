package main

import (
	. "blog-go/handler"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func main() {
	app := iris.New()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, //允许全部跨域请求
		AllowCredentials: true,
	}
	crs := cors.New(corsOptions)
	blog := app.Party("/blog", crs).AllowMethods(iris.MethodOptions)

	blog.Get("login/cookies/{name}/{value}", hero.Handler(SetCookie))
	blog.Get("login/findcookies/{name}", func(ctx iris.Context) {
		FindCookie(ctx)
	})

}
