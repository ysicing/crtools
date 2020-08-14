// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package config

import (
	"github.com/ysicing/go-utils/exfile"
	"gopkg.in/yaml.v2"
)

// Config 配置
type Config struct {
	AliKey    string `yaml:"ali_key"`
	AliSecret string `yaml:"ali_secret"`
	Region    string `yaml:"region"`
}

func exampleConfig() Config {
	return Config{
		AliKey:    "example-Shieli2r",
		AliSecret: "example-ooQue4oLohm7thaT",
		Region:    "cn-beijing",
	}
}

// WriteDefaultCfg 写默认配置文件
func WriteDefaultCfg(path string) {
	cfg, _ := yaml.Marshal(exampleConfig())
	exfile.WriteFile(path, string(cfg))
}
