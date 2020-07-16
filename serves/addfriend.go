package serves

import (
	"blog-go/structs"
	"regexp"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sparrc/go-ping"
)

//AddFriendServe -Serve 处理友链及ping信息
func AddFriendServe(friend structs.ResFriend) {
	url := friend.Siteaddress
	dropurl := DropHead(url)
	//得到ping
	ms := Pingms(dropurl)
	WriteFriend(friend, ms)
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

	ave := pinger.Statistics().AvgRtt.Milliseconds()
	println(address, " ping结果:", ave)
	return int(ave)
}

//DropHead 删除网址头部的http(s)和尾部的子路径，返回可以ping的主机地址
func DropHead(full string) string {
	r, _ := regexp.Compile("http://|https://")
	after := r.ReplaceAllString(full, "")
	r2, _ := regexp.Compile("/.*")
	after2 := r2.ReplaceAllString(after, "")
	return after2
}

//Createid 为图片生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id
}
