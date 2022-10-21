package main

import (
	"github.com/wanghaha-dev/gf/crypto/gaes"
	"github.com/wanghaha-dev/gf/os/gfile"
	"github.com/wanghaha-dev/gf/os/gres"
)

var (
	CryptoKey = []byte("x76cgqt36i9c863bzmotuf8626dxiwu0")
)

func main() {
	binContent := gfile.GetBytes("data.bin")
	binContent, err := gaes.Decrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gres.Add(binContent); err != nil {
		panic(err)
	}
	gres.Dump()
}
