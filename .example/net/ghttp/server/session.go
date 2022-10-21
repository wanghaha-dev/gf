package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
	"github.com/wanghaha-dev/gf/util/gconv"
)

func main() {
	s := g.Server()
	s.BindHandler("/session", func(r *ghttp.Request) {
		id := r.Session.GetInt("id")
		r.Session.Set("id", id+1)
		r.Response.Write("id:" + gconv.String(id))
	})
	s.SetPort(8199)
	s.Run()
}
