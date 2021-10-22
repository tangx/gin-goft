package goft

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// sysConfig 配置文件
type sysConfig struct {
	Server serverConfig `yaml:"server"`
}

// serverConfig http server 配置文件
type serverConfig struct {
	Port int32  `yaml:"port"`
	Host string `yaml:"host,omitempty"`
}

// defaultSysConfig 返回默认配置文件
func defaultSysConfig() *sysConfig {
	return &sysConfig{
		Server: serverConfig{
			Host: "",
			Port: 18089,
		},
	}
}

// loadFile 从文件中读取配置，并覆盖默认值
func (cfg *sysConfig) loadFile(cfgfile string) {
	b, err := os.ReadFile(cfgfile)
	if err != nil {
		fmt.Printf("WARNING: read file failed error: %v\n", err)
	}

	err = yaml.Unmarshal(b, config)
	if err != nil {
		fmt.Printf("WARNING: unmarshal file failed error: %v\n", err)
	}
}

var config *sysConfig

// 初始化配置文件
func init() {
	config = defaultSysConfig()
	config.loadFile("goft.yaml")
}
