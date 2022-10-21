package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"time"

	"github.com/wanghaha-dev/gf/os/gtime"
)

// 测试删除日志文件是否会重建日志文件
func main() {
	path := "/Users/john/Temp/test"
	g.Log().SetPath(path)
	for {
		g.Log().Println(gtime.Now().String())
		time.Sleep(time.Second)
	}
}
