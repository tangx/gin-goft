# 注册路由组

在 gin 中有路由组的概念， 可以理解为路由的 prefix。

## goft 增加路由组

1. 在 Goft 中增加路由组 `rg *gin.RouterGroup`

```go
type Goft struct {
	*gin.Engine
	rg *gin.RouterGroup
}
```

2. 有了路由组字段之后， 就需要使用起来。 在 `Mount()` 方法中， 增加 `group name` 传参数

```go
// Mount 参数中增加了 group 的传参
func (goft *Goft) Mount(group string, classes ...ClassController) *Goft {

	// 04.1. 注册路由组
	goft.rg = goft.Group(group)

	for _, class := range classes {
		// 03.3. 将 goft 传入到控制器中
		class.Build(goft)
	}

	return goft
}
```

有了 group name 之后， 肯定是要将其注册到 goft engine 中。  

```go
goft.rg = goft.Group(group)
```

3. 为了能在不改变控制器的情况下使用 **路由组** 路径， 需要 **重载** goft 的 `Handle` 方法。

```go
// Handle 重载 gin.Engine 的 Handle 方法。
// 04.2. 这样子路由注册的时候， 就直接挂载到了 RouterGroup 上， 有了层级关系
func (goft *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	goft.rg.Handle(httpMethod, relativePath, handlers...)

    return goft
}
```

重载 Handle 方法之后， 控制器的子路由就被路由组分组了。


## 挂载路由组

在 [main.go](/cmd/goft/main.go) 中， 为 Mount 方法增加路由组 `v1`， 并添加了一个新的路由组 `v2`

```go
	// 2. 注册路由
	g.Mount("/v1",
		classes.NewIndex(),
	)
	// 04.2. 注册多个路由组。
	g.Mount("/v2",
		classes.NewIndex(),
	)
```

启动服务后，可以看到两组路由， v1 和 v2

```bash
# cd cmd/goft/ && go run .
[GIN-debug] GET    /v1/                      --> github.com/tangx-labs/gin-goft/classes.handlerIndex (3 handlers)
[GIN-debug] GET    /v2/                      --> github.com/tangx-labs/gin-goft/classes.handlerIndex (3 handlers)
[GIN-debug] Listening and serving HTTP on :8089
```