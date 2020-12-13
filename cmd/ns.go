// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "命名空间管理",
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
		nsres := crapi.NameSpaces()
		t := table.NewWriter()
		t.AppendHeader(table.Row{"", "区域", "命名空间", "权限"})
		for id, ns := range nsres {
			t.AppendRow(table.Row{id, Region, ns.Namespace, ns.AuthorizeType})
		}
		fmt.Println(t.Render())
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
}
