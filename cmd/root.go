// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/crtools/api"
	"github.com/ysicing/crtools/config"
	"github.com/ysicing/go-utils/exfile"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	// AliKey ali key
	AliKey string
	// AliSecret ali secret
	AliSecret string
	// Region ali region
	Region  string
	cfgFile string
	// Debug debug
	Debug bool
	// Namespace ns
	Namespace string
	// Repo repo
	Repo string
	// Tailnum 数
	Tailnum int
)

var rootCmd = &cobra.Command{
	Use:   "crtools",
	Short: "阿里云镜像仓库工具",
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
	rootCmd.PersistentFlags().StringVar(&AliKey, "key", "", "aliyun accessKeyID")
	rootCmd.PersistentFlags().StringVar(&AliSecret, "secret", "", "aliyun accessSecret")
	rootCmd.PersistentFlags().StringVar(&Region, "region", "", "aliyun region (default is cn-beijing)")
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
