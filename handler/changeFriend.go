package handler

import (
	"blog-go/serves"
	"blog-go/structs"

	"github.com/kataras/iris/v12"
)

//UpdateFriend -handler 修改单个友链的信息
func UpdateFriend(ctx iris.Context) {
	var resfriend structs.ResUpdateFriend
	err := ctx.ReadJSON(&resfriend)
	if err != nil {
		println("前端传入数据不合法", err.Error())
		return
	}
	url := resfriend.Siteaddress
	dropurl := serves.DropHead(url)
	//得到ping
	ms := serves.Pingms(dropurl)
	resfriend.Ping = ms
	serves.UpdataFriendDB(resfriend)
}

//DelFriend -handler 删除某个友链
func DelFriend(ctx iris.Context) {
	id, err := ctx.GetBody()
	if err != nil {
		println("Del-前端数据不合法", err.Error())
		return
	}
	strid := string(id)
	serves.DelFriendDB(strid)
}
