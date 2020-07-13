package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

//Authority 验证root权限的中间件
func Authority(ctx iris.Context) {
	var (
		cookieNameForSessionID = "adminid"
		sess                   = sessions.New(sessions.Config{
			Cookie: cookieNameForSessionID})
	)

}
