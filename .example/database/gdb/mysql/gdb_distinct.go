package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.DB().Model("user").Distinct().CountColumn("uid,name")
}
