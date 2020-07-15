package main

import (
	. "blog-go/handler"
	_ "blog-go/serves"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.Use(Cors)
	app.Logger().SetLevel("warn")
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, //允许全部跨域请求
		AllowCredentials: true,
	}
	crs := cors.New(corsOptions)

	blog := app.Party("blog", crs).AllowMethods(iris.MethodOptions)

	blog.Post("/rootlogin", RootLogin)
	blog.Get("/rootlogout", RootLogout)
	blog.Get("/getfriends", GetFriends)
	blog.Get("/refreshms", IsRoot, Refreshms)
	blog.Post("/addfriend", IsRoot, AddFriend)
	blog.Post("/updatefriend", IsRoot, UpdateFriend)
	blog.Post("/uploadpic", IsRoot, UploadPic)
	blog.Post("/delfriend", IsRoot, DelFriend)
	blog.HandleDir("/getpics", "./uploadpics")
	app.Run(iris.Addr(":8090"))

}

func notFound(ctx iris.Context) {
	println("404-找不到此路由/路径:", ctx.RequestPath(true))
	ctx.WriteString("路由/请求地址错误")
}
