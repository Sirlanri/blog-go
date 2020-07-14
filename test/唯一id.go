package main

import uuid "github.com/satori/go.uuid"

func main7() {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	println(u1.String())
}
