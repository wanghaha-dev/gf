package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
	_ "github.com/wanghaha-dev/gf/os/gres/testdata"
)

func main() {
	g.Res().Dump()
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetPath("config-custom")
	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("config.toml")
	g.Dump(g.Config().Get("redis"))
}
