package main

import (
	"blog-go/structs"
	"crypto/md5"
	"encoding/base64"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

//Myhash 计算密码的哈希值
func Myhash(pw string) string {

	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}

//RootLogin 根用户登录
func RootLogin(ctx iris.Context) {
	var npjson structs.ResMP
	ctx.ReadJSON(&npjson)
	pwFromDB := GetPassword(npjson.Password)

	if pwFromDB.Name == "" {
		//如果数据库中没有这个邮箱
		ctx.StatusCode(403)
		ctx.WriteString("No")
	} else {
		//校验hash之后的密码
		webpw := Myhash(npjson.Password) //用户输入的密码
		if webpw == pwFromDB.Password && pwFromDB.Power == 0 {
			ctx.WriteString("yes") //密码正确并且权限为0
		} else {
			ctx.WriteString("wrong") //密码错误或不是root用户
		}
	}

}
