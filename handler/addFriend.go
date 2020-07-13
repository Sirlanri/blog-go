package handler

import (
	"blog-go/structs"

	"github.com/kataras/iris/v12"
)

//AddFriend -handler 添加一个友链
func AddFriend(ctx iris.Context) {
	var newFriend structs.ResFriend
	err := ctx.ReadJSON(&newFriend)
	if err != nil {
		println("错误-添加友链-前端传入数据错误", err.Error())
	}
}
