package serves

import (
	"blog-go/structs"
	"fmt"
	"regexp"
	"time"

	"github.com/sparrc/go-ping"
)

//AddFriend -Serve 处理友链及ping信息
func AddFriend(friend structs.Friend) {
	url := friend.URL
	dropurl := dropHead(url)
	//得到ping
	ms := pingms(dropurl)

}

//pingms 传入主机地址，返回ping值
func pingms(address string) (result int) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		println(err.Error())
		//ping不通，默认服务器挂了
		result = 1000
	}
	//true需要在管理员下运行
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Second * 20
	pinger.Count = 4
	pinger.Interval = time.Millisecond * 10
	pinger.Run()

	rev := pinger.PacketsRecv
	//全部丢包，就默认服务器挂了
	if rev == 0 {
		result = 1000
	}
	sta := pinger.Statistics().Rtts
	totalms := sta[0] + sta[1] + sta[2] + sta[3]
	totalms.Seconds()
	result = int(totalms.Milliseconds() / int64(rev))
	fmt.Print(result)
	return
}

//dropHead 删除网址头部的http，返回可以ping的主机地址
func dropHead(full string) string {
	r, _ := regexp.Compile("http://|https://")
	after := r.ReplaceAllString(full, "")
	return after
}
