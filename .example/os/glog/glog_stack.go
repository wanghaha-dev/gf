package main

import (
	"fmt"
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.Log().PrintStack()

	fmt.Println(g.Log().GetStack())
}
