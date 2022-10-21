package main

import (
	"github.com/wanghaha-dev/gf/os/gfile"
	"github.com/wanghaha-dev/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
