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

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "镜像仓库搜索",
	Aliases: []string{"sn"},
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
		searchres := crapi.PreSearch(SearchName)
		t := table.NewWriter()
		t.AppendHeader(table.Row{"", "区域", "镜像", "镜像更新时间"})
		for id, repo := range searchres {
			t.AppendRow(table.Row{id, Region, fmt.Sprintf("registry.%v.aliyuncs.com/%v/%v:%v", Region, repo.RepoNamespace, repo.RepoName, repo.LastTag),
				extime.UnixInt642String(repo.ImageUpdate / 1000)})
		}
		fmt.Println(t.Render())
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().StringVar(&Namespace, "ns", "", "命名空间")
	searchCmd.PersistentFlags().StringVar(&SearchName, "sn", "", "搜索字符串")
}
