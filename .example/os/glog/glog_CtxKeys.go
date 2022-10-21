package main

import (
	"context"
	"github.com/wanghaha-dev/gf/frame/g"
)

func main() {
	g.Log().SetCtxKeys("TraceId", "SpanId", "Test")
	ctx := context.WithValue(context.Background(), "TraceId", "1234567890")
	ctx = context.WithValue(ctx, "SpanId", "abcdefg")

	g.Log().Ctx(ctx).Print(1, 2, 3)
}
