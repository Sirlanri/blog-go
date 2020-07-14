package main

import "regexp"

func main5() {
	source := "http://ri-co.cn"
	r, _ := regexp.Compile("http://|https://")
	after := r.ReplaceAllString(source, "")
	println(after)
}
