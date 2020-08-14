// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import "k8s.io/klog"

// LogDebug debug
func LogDebug(msg interface{}, mode bool) {
	if mode {
		klog.Info(msg)
	}
}
