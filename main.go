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

	app.Logger().SetLevel("warn")
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})

	blog := app.Party("blog", crs).AllowMethods(iris.MethodOptions)
	{
		blog.Post("/rootlogin", RootLogin)
		blog.Get("/rootlogout", RootLogout)
		blog.Get("/getfriends", GetFriends)
		blog.Get("/refreshms", IsRoot, Refreshms)
		blog.Post("/addfriend", IsRoot, AddFriend)
		blog.Post("/updatefriend", IsRoot, UpdateFriend)
		blog.Post("/uploadpic", IsRoot, UploadPic)
		blog.Post("/delfriend", IsRoot, DelFriend)
		blog.HandleDir("/getpics", iris.Dir("./uploadpics"))
	}

	app.Run(iris.Addr(":9004"))

}

func notFound(ctx iris.Context) {
	println("404-找不到此路由/路径:", ctx.RequestPath(true))
	ctx.WriteString("路由/请求地址错误")
}
