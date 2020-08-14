// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
	"k8s.io/klog"
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
		for _, ns := range nsres {
			klog.Infof("ns: %v, authorize: %v\n", ns.Namespace, ns.AuthorizeType)
		}
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
