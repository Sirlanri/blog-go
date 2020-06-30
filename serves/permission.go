package serves

import (
	"blog-go/structs"
	"crypto/md5"
	"fmt"

	"github.com/kataras/iris/v12"
)

func rootLogin(ctx iris.Context) {
	var npjson structs.NamePw
	npjson.Name = "lan"
	npjson.Password = "cooooooo"
	ctx.ReadJSON(&npjson)
	ha := md5.New().Sum([]byte(npjson.Password))
	fmt.Println(ha)

}
