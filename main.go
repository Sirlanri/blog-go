package main

import (
	. "blog-go/handler"
	. "blog-go/serves"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func main() {
	app := iris.New()

	blog := app.Party("/blog")
	blog.Get("login/cookies/{name}/{value}", hero.Handler(SetCookie))
	blog.Get("login/findcookies/{name}", func(ctx iris.Context) {
		FindCookie(ctx)
	})

}
