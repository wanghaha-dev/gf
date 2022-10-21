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
	binContent, err := gres.Pack("public,config")
	if err != nil {
		panic(err)
	}
	binContent, err = gaes.Encrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gfile.PutBytes("data.bin", binContent); err != nil {
		panic(err)
	}
}
