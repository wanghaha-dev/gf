package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/os/glog"
)

func main() {
	g.Log().SetLevelPrefix(glog.LEVEL_DEBU, "debug")
	g.Log().Debug("test")
}
