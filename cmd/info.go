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

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "信息",
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
	rootCmd.AddCommand(infoCmd)
}
