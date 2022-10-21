package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		if r.GetInt("type") == 1 {
			r.Response.Writeln("john")
		}
		r.Response.Writeln("smith")
	})
	s.SetPort(8199)
	s.Run()
}
