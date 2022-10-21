package main

import (
	"context"
	"fmt"
	"github.com/wanghaha-dev/gf/errors/gerror"
	"github.com/wanghaha-dev/gf/frame/g"
	"github.com/wanghaha-dev/gf/util/gmeta"
)

type UserCreateReq struct {
	gmeta.Meta `v:"UserCreateReq"`
	Name       string
	Pass       string
}

func UserCreateReqChecker(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
	user := &UserCreateReq{}
	if v, ok := data.(*UserCreateReq); ok {
		user = v
	}
	// SELECT COUNT(*) FROM `user` WHERE `name` = xxx
	count, err := g.Model("user").Ctx(ctx).Where("name", user.Name).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.Newf(`The name "%s" is already token`, user.Name)
	}
	return nil
}

func main() {
	user := &UserCreateReq{
		Name: "john",
		Pass: "123456",
	}
	err := g.Validator().RuleFunc("UserCreateReq", UserCreateReqChecker).CheckStruct(user)
	fmt.Println(err)
}
