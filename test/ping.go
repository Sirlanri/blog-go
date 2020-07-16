package main

import (
	"regexp"
	"time"

	"github.com/sparrc/go-ping"
)

func main() {
	website := "https://shawnzhou.world/tags/%E5%BC%BA%E8%BF%9E%E9%80%9A%E5%88%86%E9%87%8F/"
	after := DropHead(website)
	println(Pingms(after))
	return
}

//Pingms 传入主机地址，返回ping值
func Pingms(address string) (result int) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		println(err.Error())
		//域名无法解析（比如被墙），默认服务器挂了
		result = 1000
		return
	}
	//true需要在管理员下运行
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Second * 20
	pinger.Count = 4
	pinger.Interval = time.Millisecond * 10
	pinger.Run()

	//rev := pinger.PacketsRecv

	//sta := pinger.Statistics().Rtts
	ave := pinger.Statistics().AvgRtt.Milliseconds()

	return int(ave)
}

//DropHead 删除网址头部的http(s)和尾部的子路径，返回可以ping的主机地址
func DropHead(full string) string {
	r, _ := regexp.Compile("http://|https://")
	after := r.ReplaceAllString(full, "")
	r2, _ := regexp.Compile("/.*")
	after2 := r2.ReplaceAllString(after, "")
	println(after2)
	return after2
}
