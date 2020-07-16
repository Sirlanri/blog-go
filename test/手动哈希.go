package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func main3() {
	str := ""
	fmt.Println(Myhash(str))
}

//Myhash 计算密码的哈希值
func Myhash(pw string) string {
	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}
