package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"time"

	"github.com/wanghaha-dev/gf/os/gtime"
	"github.com/wanghaha-dev/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		g.Log().SetDebug(false)
	})
	for {
		g.Log().Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
