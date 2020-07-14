package handler

import (
	. "blog-go/serves"
	"blog-go/structs"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieName = "rootid"
	rootsess   = sessions.New(sessions.Config{
		Cookie: cookieName})
)

//RootLogin -handler root用户登录的
func RootLogin(ctx iris.Context) {
	println("执行登录程序")
	session := rootsess.Start(ctx)

	//执行验证
	var npjson structs.ResMP
	ctx.ReadJSON(&npjson)
	result := RootConfirm(npjson)
	if result == 0 {
		ctx.WriteString("yes")
	}
	if result == 1 {
		//如果数据库中没有这个邮箱
		ctx.WriteString("no")
		return
	}
	if result == 2 {
		ctx.WriteString("wrong")
		return
	}

	//设置验证状态root为true
	session.Set("root", true)
	test, err := session.GetBoolean("root")
	if err != nil {
		println(err.Error())
	}
	println(test)
	println("root用户登录，授予权限")
}

//RootLogout -handler root用户退出登录，完成后返回'done'
func RootLogout(ctx iris.Context) {
	println("注销用户登录状态")
	session := rootsess.Start(ctx)
	//撤销权限
	session.Set("root", false)
	//删除session
	rootsess.DestroyByID("rootid")
	//注销成功，返回done
	ctx.WriteString("done")
}

//IsRoot -handler 判断用户是否为root的中间件
func IsRoot(ctx iris.Context) {
	thissess := rootsess.Start(ctx)
	auth, err := thissess.GetBoolean("root")

	if err != nil {
		println("权限判断出错", err.Error())
	}
	if !auth {
		ctx.StatusCode(403)
		return
	}
	//权限通过，执行下一个handler
	ctx.Next()
	return

}
