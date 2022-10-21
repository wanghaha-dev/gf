package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.RedirectTo("/login")
	})
	s.BindHandler("/login", func(r *ghttp.Request) {
		r.Response.Writeln("Login First")
	})
	s.SetPort(8199)
	s.Run()
}
