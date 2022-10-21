package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
)

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.tpl", g.Map{
			"title": "Test",
			"name":  "John",
			"score": 100,
		})
	})
	s.SetPort(8199)
	s.Run()
}
