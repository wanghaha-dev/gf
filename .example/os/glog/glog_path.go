package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/os/gfile"
)

// 设置日志输出路径
func main() {
	path := "/tmp/glog"
	g.Log().SetPath(path)
	g.Log().Println("日志内容")
	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)
}
