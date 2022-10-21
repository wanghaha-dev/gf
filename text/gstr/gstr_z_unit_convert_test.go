// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/wanghaha-dev/gf.

// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/wanghaha-dev/gf/test/gtest"
	"github.com/wanghaha-dev/gf/text/gstr"
)

func Test_OctStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.OctStr(`\346\200\241`), "怡")
	})
}
