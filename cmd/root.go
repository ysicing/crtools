// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
	"github.com/ysicing/crtools/config"
	"github.com/ysicing/ext/utils/exfile"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	// Key ali key
	Key string
	// Secret ali secret
	Secret string
	// Region ali region
	Region string
	// Provider 服务商
	Provider string
	cfgFile  string
	// Debug debug
	Debug bool
	// Namespace ns
	Namespace string
	// Repo repo
	Repo string
	// Tailnum 数
	Tailnum int
	// SearchName sn
	SearchName string
)

var rootCmd = &cobra.Command{
	Use:   "crtools",
	Short: "镜像仓库工具",
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cr.yaml)")
	rootCmd.PersistentFlags().StringVar(&Key, "key", "", "cloud accessKeyID")
	rootCmd.PersistentFlags().StringVar(&Secret, "secret", "", "cloud accessSecret")
	rootCmd.PersistentFlags().StringVar(&Region, "region", "cn-beijing", "cloud region")
	rootCmd.PersistentFlags().StringVar(&Provider, "provider", "aliyun", "cloud provider")
	rootCmd.PersistentFlags().BoolVar(&api.Debug, "debug", false, "debug mode")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.DisableSuggestions = false
}

func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			os.Exit(1)
		}
		cfgFile = fmt.Sprintf("%v/%v", home, ".cr.yaml")
	}
	if !exfile.CheckFileExistsv2(cfgFile) {
		config.WriteDefaultCfg(cfgFile)
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
