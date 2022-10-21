package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/os/glog"
)

// 设置日志等级，过滤掉Info日志信息
func main() {
	g.Log().Info("info1")
	g.Log().SetLevel(glog.LEVEL_ALL ^ glog.LEVEL_INFO)
	g.Log().Info("info2")
}
