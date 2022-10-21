// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/wanghaha-dev/gf.

package gcron_test

import (
	"time"

	"github.com/wanghaha-dev/gf/os/gcron"
	"github.com/wanghaha-dev/gf/os/glog"
)

func Example_cronAddSingleton() {
	gcron.AddSingleton("* * * * * *", func() {
		glog.Println("doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
