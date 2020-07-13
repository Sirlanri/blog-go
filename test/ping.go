package main

import (
	"fmt"
	"time"

	"github.com/sparrc/go-ping"
)

func main4() {
	pinger, err := ping.NewPinger("liangyj_blog.gitee.io")
	if err != nil {
		println(err.Error())
	}
	//true需要在管理员下运行
	pinger.SetPrivileged(true)
	pinger.Timeout = time.Second * 20
	pinger.Count = 4
	pinger.Interval = time.Millisecond * 10
	pinger.Run()

	rev := pinger.PacketsRecv
	var result int
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
