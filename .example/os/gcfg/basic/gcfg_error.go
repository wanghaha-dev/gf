package main

import (
	"fmt"

	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	fmt.Println(g.Config().Get("none"))
}
