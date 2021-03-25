// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/crtools/cmd"
	"github.com/ysicing/ext/logger/zlog"
)

func init()  {
	cfg := zlog.Config{
		Simple:      true,
		HookFunc:    nil,
		WriteLog:    false,
		WriteJSON:   false,
		ServiceName: "crtools",
	}
	zlog.InitZlog(&cfg)
}

func main() {
	cmd.Execute()
}
