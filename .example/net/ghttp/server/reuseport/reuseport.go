package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
	"github.com/wanghaha-dev/gf/os/glog"
)

func main() {
	s1 := ghttp.GetServer("s1")
	s1.SetPort(8882)
	s1.BindHandler("/", func(r *ghttp.Request) {
		glog.Println("s1")
		r.Response.Writeln("s1")
	})
	s1.Start()

	s2 := ghttp.GetServer("s2")
	s2.SetPort(8882)
	s2.BindHandler("/", func(r *ghttp.Request) {
		glog.Println("s2")
		r.Response.Writeln("s2")
	})
	s2.Start()

	g.Wait()
}
