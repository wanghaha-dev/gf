package main

import (
	"fmt"

	"github.com/wanghaha-dev/gf/encoding/gcompress"
	"github.com/wanghaha-dev/gf/os/gfile"
)

func main() {
	err := gcompress.UnZipContent(
		gfile.GetBytes(`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg\encoding\gcompress\data.zip`),
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg`,
	)
	fmt.Println(err)
}
