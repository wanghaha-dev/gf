package main

import (
	"context"
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/util/gconv"
	"github.com/wanghaha-dev/gf/util/gvalid"
)

func main() {
	type User struct {
		Name string `v:"required#请输入用户姓名"`
		Type int    `v:"required#请选择用户类型"`
	}
	data := g.Map{
		"name": "john",
	}
	user := User{}
	if err := gconv.Scan(data, &user); err != nil {
		panic(err)
	}
	err := gvalid.CheckStructWithData(context.TODO(), user, data, nil)
	// 也可以使用
	// err := g.Validator().Data(data).CheckStruct(user)
	if err != nil {
		g.Dump(err.Items())
	}
}
