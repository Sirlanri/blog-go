package main

import (
	. "blog-go/handler"
	_ "blog-go/serves"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, //允许全部跨域请求
		AllowCredentials: true,
	}
	crs := cors.New(corsOptions)
	blog := app.Party("blog", crs).AllowMethods(iris.MethodOptions)

	blog.Post("/rootlogin", RootLogin)
	blog.Get("/rootlogout", RootLogout)
	blog.Get("/getfriends", GetFriends)
	blog.Get("/refreshms", Refreshms)
	blog.Post("/addfriend", IsRoot, AddFriend)
	blog.Post("/uploadpic", IsRoot, UploadPic)
	blog.HandleDir("/getpics", "./uploadpics")
	app.Run(iris.Addr(":8090"))

}

func notFound(ctx iris.Context) {
	println("404-找不到此路由/路径:", ctx.RequestPath(true))
	ctx.WriteString("路由/请求地址错误")
}
