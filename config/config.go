// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package config

import (
	"github.com/ysicing/ext/utils/exfile"
	"gopkg.in/yaml.v2"
)

// Config 配置
type Config struct {
	Ali    CloudConfig `yaml:"ali"`
	Ucloud CloudConfig `yaml:"ucloud"`
}

// CloudConfig 配置
type CloudConfig struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
	Region string `yaml:"region"`
}

func exampleConfig() Config {
	return Config{
		Ali: CloudConfig{
			Key:    "meemohBelawoh5oh",
			Secret: "Thio8dahth6eig2ung2gohreiphie7ge",
			Region: "cn-beijing",
		},
		Ucloud: CloudConfig{
			Key:    "meemohBelawoh5oh",
			Secret: "Thio8dahth6eig2ung2gohreiphie7ge",
			Region: "cn-bj2-04",
		},
	}
}

// WriteDefaultCfg 写默认配置文件
func WriteDefaultCfg(path string) {
	cfg, _ := yaml.Marshal(exampleConfig())
	exfile.WriteFile(path, string(cfg))
}
