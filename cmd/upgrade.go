// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangle201210/githubapi/repos"
	"github.com/ysicing/go-utils/excmd"
	"runtime"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade crtools",
	Run: func(cmd *cobra.Command, args []string) {
		var pkg = repos.Pkg{"ysicing", "crtools"}
		lastag, _ := pkg.LastTag()
		if lastag.Name != Version {
			if runtime.GOOS != "linux" {
				excmd.RunCmd("/bin/zsh", "-c", "brew install crtools")
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
