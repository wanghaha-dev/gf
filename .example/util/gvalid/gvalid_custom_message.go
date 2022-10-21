package main

import (
	"context"
	"fmt"
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check(context.TODO(), "", "required", nil)
	fmt.Println(err.String())
}
