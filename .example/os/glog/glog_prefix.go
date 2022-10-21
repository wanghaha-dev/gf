package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.Log().SetPrefix("[API]")
	g.Log().Println("hello world")
	g.Log().Error("error occurred")
}
