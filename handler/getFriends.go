package handler

import (
	"blog-go/serves"

	"github.com/kataras/iris/v12"
)

//GetFriends Get-handler 返回友链json，格式为'list':[Friends]
//不需要验证权限
func GetFriends(ctx iris.Context) {
	friends := serves.GetFriendDB()

	list1 := map[string]interface{}{
		"list": friends,
	}
	ctx.JSON(list1)

}
