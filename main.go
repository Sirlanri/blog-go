package main

import (
	. "blog-go/handler"
	_ "blog-go/serves"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, //允许全部跨域请求
		AllowCredentials: true,
	}
	crs := cors.New(corsOptions)
	blog := app.Party("/blog", crs).AllowMethods(iris.MethodOptions)

	blog.Post("rootlogin", RootLogin)
	blog.Post("newfriend", NewFriend)
	app.Run(iris.Addr(":8090"))

}
