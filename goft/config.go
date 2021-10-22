package goft

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// goftConfig 配置文件
type goftConfig struct {
	Server serverConfig `yaml:"server"`
}

// serverConfig http server 配置文件
type serverConfig struct {
	Port int32  `yaml:"port"`
	Host string `yaml:"host"`
}

// defaultGoftConfig 返回默认配置文件
func defaultGoftConfig() *goftConfig {
	return &goftConfig{
		Server: serverConfig{
			Host: "",
			Port: 18089,
		},
	}
}

// loadFile 从文件中读取配置，并覆盖默认值
func (cfg *goftConfig) loadFile(cfgfiles ...string) {
	for _, cfgfile := range cfgfiles {
		b, err := os.ReadFile(cfgfile)
		if err != nil {
			fmt.Printf("WARNING: read file failed error: %v\n", err)
		}

		err = yaml.Unmarshal(b, config)
		if err != nil {
			fmt.Printf("WARNING: unmarshal file failed error: %v\n", err)
		}
	}
}

func (cfg *goftConfig) dump() {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}

	err = os.WriteFile("goft.default.yml", data, 0644)
	if err != nil {
		log.Printf("WARNING: write default config failed: %v\n", err)
	}
}

var config *goftConfig

// 初始化配置文件
func init() {
	config = defaultGoftConfig()
	config.dump()
	config.loadFile("goft.default.yml", "goft.yml")
}
