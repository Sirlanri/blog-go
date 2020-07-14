package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "adminid"
	sess                   = sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	//验证权限
	sessionNow := sess.Start(ctx)
	fmt.Println(sessionNow)

	auth, err := sess.Start(ctx).GetBoolean("root")
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
	session.Set("root", true)
	println(session)
	ctx.WriteString("成功登录")
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)
	//撤销验证
	session.Set("root", false)
	//删除session
	sess.DestroyByID("adminid")
	ctx.WriteString("已退出")
}

type loginInfor struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func loginp(ctx iris.Context) {
	session1 := sess.Start(ctx)
	var infor loginInfor
	if err := ctx.ReadJSON(&infor); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("非法请求格式")
		fmt.Println(err.Error())
		return
	}
	if infor.Name == "rico" && infor.Password == "123" {
		session1.Set("root", true)
		ctx.WriteString("成功登录")
		return
	}
	ctx.WriteString("验证失败")

}

//顺便保存点内容
func flashTest(ctx iris.Context) {
	s := sessions.Get(ctx)

	s.SetFlash("name", "Rico")
	ctx.WriteString("写入名字已完成~")
}
func getFlash(ctx iris.Context) {
	s := sessions.Get(ctx)
	name := s.GetFlashString("name")
	if name == "" {
		ctx.WriteString("name为空")
	} else {
		ctx.Writef("你好鸭%s", name)
	}

}

func main() {
	app := iris.New()
	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Post("/loginp", loginp)
	app.Get("/setflash", flashTest)
	app.Get("/getflash", getFlash)
	app.Run(iris.Addr(":8090"))
}
