package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.Log().Debug(g.Map{"uid": 100, "name": "john"})

	type User struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	g.Log().Debug(User{100, "john"})
}
