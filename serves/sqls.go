package main

import (
	"blog-go/structs"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	println("自动初始化数据库函数")
	ConnectDB()
}

//ConnectDB 初始化时，连接数据库
func ConnectDB() *sql.DB {
	db, _ = sql.Open("mysql", "root:123456@/blog")

	if db.Ping() != nil {
		println("初始化-数据库-用户/密码/库验证失败", db.Ping().Error())
		return nil
	}
	return db

}

//GetPassword 传入邮箱,从数据库中获取名字和密码
func GetPassword(mail string) structs.Gotnp {
	var npjson structs.Gotnp
	row := db.QueryRow("select name,password,authority from users where mail=?", mail)
	row.Scan(&npjson.Name, &npjson.Password, &npjson.Power)

	return npjson

}

func main() {
	ConnectDB()
	GetPassword("mail@ri-c.cn")
}
