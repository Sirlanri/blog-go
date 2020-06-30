package serves

import (
	"blog-go/structs"
	"crypto/md5"
	"encoding/base64"

	_ "github.com/go-sql-driver/mysql"
)

//Myhash 计算密码的哈希值
func Myhash(pw string) string {

	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}

//RootConfirm 根用户登录
func RootConfirm(mailAndPw structs.ResMP) int {
	pwFromDB := GetPassword(mailAndPw.Password)

	if pwFromDB.Name == "" {
		//如果数据库中没有这个邮箱
		return 1
	}

	//校验hash之后的密码
	webpw := Myhash(mailAndPw.Password) //用户输入的密码
	if webpw == pwFromDB.Password && pwFromDB.Power == 0 {
		return 0 //密码正确并且权限为0
	}
	return 2 //密码错误或不是root用户

}
