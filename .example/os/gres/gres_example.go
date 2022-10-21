package main

import (
	_ "github.com/wanghaha-dev/gf/os/gres/testdata/example/boot"

	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/template", func(r *ghttp.Request) {
			r.Response.WriteTplDefault(g.Map{
				"name": "GoFrame",
			})
		})
	})
	s.Run()
}
