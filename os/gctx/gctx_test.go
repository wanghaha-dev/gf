// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/wanghaha-dev/gf.

package gctx_test

import (
	"context"
	"github.com/wanghaha-dev/gf/os/gctx"
	"github.com/wanghaha-dev/gf/text/gstr"
	"testing"

	"github.com/wanghaha-dev/gf/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := gctx.New()
		t.AssertNE(ctx, nil)
		t.AssertNE(gctx.Value(ctx), "")
	})
}

func Test_WithCtx(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.WithCtx(ctx)
		t.AssertNE(gctx.Value(ctx), "")
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_WithPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.WithPrefix(ctx, "H-")
		t.Assert(gstr.Contains(gctx.Value(ctx), "H-"), true)
		t.Assert(ctx.Value("TEST"), 1)
	})
}

func Test_WithValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.WithValue(context.TODO(), "TEST", 1)
		ctx = gctx.WithValue(ctx, "123")
		t.Assert(gctx.Value(ctx), "123")
		t.Assert(ctx.Value("TEST"), 1)
	})
}
