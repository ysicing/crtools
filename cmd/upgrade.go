// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/wangle201210/githubapi/repos"
	"github.com/ysicing/ext/utils/excmd"
	"k8s.io/klog"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade crtools",
	Aliases: []string{"up"},
	Run: func(cmd *cobra.Command, args []string) {
		var pkg = repos.Pkg{"ysicing", "crtools"}
		lastag, _ := pkg.LastTag()
		if lastag.Name != Version {
			if runtime.GOOS != "linux" {
				klog.Info(excmd.RunCmdRes("/bin/zsh", "-c", "brew upgrade ysicing/tap/crtools"))
			} else {
				newbin := fmt.Sprintf("https://github.com/ysicing/crtools/releases/download/%v/crtools_linux_amd64", lastag.Name)
				excmd.DownloadFile(newbin, "/usr/local/bin/crtools")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
