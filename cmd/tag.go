// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
	"github.com/ysicing/ext/ztime"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "镜像标签管理",
	Run: func(cmd *cobra.Command, args []string) {
		if len(Key) == 0 {
			Key = viper.GetString("ali_key")
		} else {
			viper.Set("ali_key", Key)
		}
		if len(Secret) == 0 {
			Secret = viper.GetString("ali_secret")
		} else {
			viper.Set("ali_secret", Secret)
		}
		if len(Region) == 0 {
			Region = viper.GetString("region")
		} else {
			viper.Set("region", Region)
		}
		crapi := api.NewAPI(Key, Secret, Region)
		tagsres := crapi.Tags(Namespace, Repo, Tailnum)
		t := table.NewWriter()
		t.AppendHeader(table.Row{"", "镜像", "版本", "镜像大小", "更新时间"})
		for id, tag := range tagsres {
			t.AppendRow(table.Row{id, fmt.Sprintf("registry.%v.aliyuncs.com/%v/%v:%v", Region, Namespace, Repo, tag.Tag), tag.Tag,
				tag.ImageSize, ztime.UnixInt642String(tag.ImageUpdate / 1000)})
		}
		fmt.Println(t.Render())
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.PersistentFlags().StringVar(&Namespace, "ns", "", "命名空间")
	tagsCmd.PersistentFlags().StringVar(&Repo, "repo", "", "镜像名")
	tagsCmd.PersistentFlags().IntVar(&Tailnum, "tail", 1, "显示数目(默认最后一个)")
}
