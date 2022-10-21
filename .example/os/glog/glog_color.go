package main

import (
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.Log().Print("Print")
	g.Log().Debug("Debug")
	g.Log().Info("Info")
	g.Log().Notice("Notice")
	g.Log().Warning("Warning")
	g.Log().Error("Error")
	g.Log().Critical("Critical")
}
