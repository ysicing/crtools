// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
	"github.com/ysicing/go-utils/extime"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "镜像仓库管理",
	//PreRun: func(cmd *cobra.Command, args []string) {
	//	if len(Namespace) == 0 {
	//		klog.Exit("命令空间不能为空")
	//	}
	//},
	Run: func(cmd *cobra.Command, args []string) {
		if len(AliKey) == 0 {
			AliKey = viper.GetString("ali_key")
		} else {
			viper.Set("ali_key", AliKey)
		}
		if len(AliSecret) == 0 {
			AliSecret = viper.GetString("ali_secret")
		} else {
			viper.Set("ali_secret", AliSecret)
		}
		if len(Region) == 0 {
			Region = viper.GetString("region")
		} else {
			viper.Set("region", Region)
		}
		crapi := api.NewAPI(AliKey, AliSecret, Region)
		repores := crapi.Repos(Tailnum, Namespace)
		t := table.NewWriter()
		t.AppendHeader(table.Row{"", "区域", "镜像", "更新时间"})
		for id, repo := range repores {
			t.AppendRow(table.Row{id, Region, fmt.Sprintf("registry.%v.aliyuncs.com/%v/%v:%v", Region, repo.RepoNamespace, repo.RepoName, repo.LastTag),
				extime.UnixInt642String(repo.GmtModified / 1000)})
		}
		fmt.Println(t.Render())
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.PersistentFlags().StringVar(&Namespace, "ns", "", "命名空间")
	repoCmd.PersistentFlags().IntVar(&Tailnum, "tail", 10, "显示数目")
}
