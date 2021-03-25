// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import (
	"github.com/ysicing/ext/logger/zlog"
)

// LogDebug debug
func LogDebug(msg interface{}, mode bool) {
	if mode {
		zlog.Info("[DEBUG]", msg)
	}
}
