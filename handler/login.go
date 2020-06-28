package handler

import "github.com/kataras/iris/v12"

//SetCookie 获得Cookie请求
func SetCookie(ctx iris.Context) {
	println("接收到cookie请求")
	name := ctx.Params().Get("name")
	value := ctx.Params().Get("value")
	ctx.SetCookieKV(name, value)
	ctx.Request().Cookie(name)
	ctx.Writef("cookie值是%s", value)
}

//SetSesson 设置session
func SetSesson(ctx iris.Context) {

}

//FindCookie 查找cookie
func FindCookie(ctx iris.Context) {
	name := ctx.Params().Get("name")
	value := ctx.GetCookie(name) //检索获取cookie
	ctx.WriteString(value)
}

//Ifpermission 中间件，判断是否已经登录获得权限
func Ifpermission(ctx iris.Context) {

}
