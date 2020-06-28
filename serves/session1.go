package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	//验证权限
	sessionNow := sess.Start(ctx)
	fmt.Println(sessionNow)

	auth, err := sess.Start(ctx).GetBoolean("authenticated")
	if err != nil {
		fmt.Println(err.Error())
	}

	if !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	ctx.WriteString("成功登录啦")
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)
	fmt.Println(session)
	//执行验证

	//设置验证状态true
	session.Set("authenticated", true)
	println(session)
	ctx.WriteString("成功登录")
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)
	//撤销验证
	session.Set("authenticated", false)
	ctx.WriteString("已退出")
}

func main() {
	app := iris.New()
	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Run(iris.Addr(":8090"))
}
