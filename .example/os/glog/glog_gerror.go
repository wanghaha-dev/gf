package main

import (
	"errors"
	"github.com/wanghaha-dev/gf/frame/g"

	"github.com/wanghaha-dev/gf/errors/gerror"
)

func MakeError() error {
	return errors.New("connection closed with normal error")
}

func MakeGError() error {
	return gerror.New("connection closed with gerror")
}

func TestGError() {
	err1 := MakeError()
	err2 := MakeGError()
	g.Log().Error(err1)
	g.Log().Error(err2)
}

func main() {
	TestGError()
}
