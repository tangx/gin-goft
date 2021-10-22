# 读取配置文件

为了使启动参数配置更加可视化和集中化， 可以引入 **配置文件** 作为载体。在没有配置文件的时候使用默认值， 在有配置文件的时候使用新值覆盖默认值。


[config.go](/goft/config.go) 完成了所有工作。

定义配置文件。 goftConfig 为所有配置的总纲， 其中嵌套的字段为不同模块的配置。 例如这里的 `Server` 为 goft httpserver 启动的参数

```go
// goftConfig 配置文件
type goftConfig struct {
	Server serverConfig `yaml:"server"`
}

// serverConfig httpserver 配置文件
type serverConfig struct {
	Port int32  `yaml:"port"`
	Host string `yaml:"host"`
}
```

创建默认 goftConfig 时， 其中定义了一些默认参数。

```go
// defaultGoftConfig 返回默认配置文件
func defaultGoftConfig() *sysConfig {
	return &sysConfig{
		Server: serverConfig{
			Host: "",
			Port: 18089,
		},
	}
}
```

### 默认配置与自定义配置

虽然说可以在配置文件中进行自定义配置， 作为使用者可能并不知道本身有什么参数。 因此使用 `dump()` 方法随时随地持久化一份最新的。

在通过 `loadFile()` 方法， 传入自定义配置值。

```go
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


// dump 持久化默认配置
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
```

### 初始化配置

在 `config.go` 文件中创建一个全局的 `config` 对象。 并在 `init` 初始化的时候进行 **默认值** 赋值与 **配置文件** 赋值。

```go
var config *goftConfig

// 初始化配置文件
func init() {
	config = defaultGoftConfig()
	config.loadFile("goft.yaml")
}
```

## 配置文件 demo

在定义配置文件的时候， 可以只写需要修改的配置。

```yaml
server:
  host: 192.168.100.100
#   port: 8089
```