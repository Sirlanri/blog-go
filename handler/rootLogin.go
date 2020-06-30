package handler

import (
	. "blog-go/serves"
	"blog-go/structs"

	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieName = "adminid"
	rootsess   = sessions.New(sessions.Config{
		Cookie: cookieName})
)

func RootLogin(ctx iris.Context) {
	session := rootsess.Start(ctx)
	fmt.Println(session)

	//执行验证
	var npjson structs.ResMP
	ctx.ReadJSON(&npjson)
	result := RootConfirm(npjson)
	if result == 0 {
		ctx.WriteString("yes")
		return
	}
	if result == 1 {
		//如果数据库中没有这个邮箱
		ctx.WriteString("No")
		return
	}
	if result == 2 {
		ctx.WriteString("wrong")
		return
	}

	//设置验证状态true
	session.Set("authenticated", true)
	println(session)
	ctx.WriteString("成功登录")
}
