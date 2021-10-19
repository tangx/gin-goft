# 级联路由组挂载注册

在上一篇中实现了路由的分组挂载， 但是遗留了一个问题： 丢失了 gin 中的 **路由级联注册** 的特性。

这一篇就找回来。


## GoftGroup

首先，需要对原来的 `gin.RouterGroup` 进行一定的扩展。 在 [goft_group.go](/goft/goft_group.go) 中我们自己封装一个 `GoftGroup`， 在其中 **匿名嵌套** `*gin.RouterGroup`

```go
type GoftGroup struct {
	*gin.RouterGroup
}
```

在 gin 中， engine 的 `Handle()` 也是直接调用了嵌套的 `gin.RouterGroup` 的 `Handle` 方法。 
因此， goft 的 `Mount()` 也同样下放到 `GoftGroup` 的 `Mount()` 中

```go
// Mount 在 GoftGroup 上绑定/注册 控制器
func (gg *GoftGroup) Mount(group string, claess ...ClassController) *GoftGroup {
	grp := newGoftGroup(gg, group)

	for _, class := range claess {
		class.Build(grp)
	}

	return grp
}
```

与之对应的， `ClassController` 的接口也需要做相应的调整。

[class_controller.go](/goft/class_controller.go)

```go
type ClassController interface {
	// Build(goft *Goft)  // 旧的
	Build(goft *GoftGroup)
}
```

在实现的 `ClassController` 接口的 **控制器** 中也需要进行响应调整。

[index.go](/cmd/goft/classes/index.go#L20)

```go
func (index *Index) Build(goft *goft.GoftGroup) {
	goft.Handle("GET", "/index", handlerIndex)
}
```


## Goft 改造

首先， Goft 不能在直接使用 `gin.RouterGroup` 了， 而是使用封装之后的 `GoftGroup`

```go
type Goft struct {
	*gin.Engine
	// rg *gin.RouterGroup
	gg *GoftGroup
}
```

### Mount

对于控制器中路由的挂载， 就直接下沉给 `GoftGroup` 执行。 为了下沉的时候遇到 GoftGroup 为 nil 发生 panic 的情况， 做了一个安全保护。

```go
// Mount 挂载控制器
// 03.1. 关联控制器与 goft
// 03.2. 返回 *GoftGroup 是为了方便链式调用
func (goft *Goft) Mount(group string, classes ...ClassController) *GoftGroup {

	// 04.1. 注册路由组
	if goft.gg == nil {
		goft.gg = baseGoftGroup(goft, "/")
	}

	return goft.gg.Mount(group, classes...)
}
```

### BasePath

在将服务上线到 k8s 中， 会使用使用 ingress 进行请求转发， 这个时候通常使用 **服务名称** 作为 uri 的 **第一段** 进行转发匹配。 例如

```
http://127.0.0.1/demo/v1/api
```

为了更好的兼容这种情况， 在 Goft 中增加了 `BasePath` 方法， 以设置 uri 的 prefix。

```go
// BasePath 设置 Goft 的根路由
func (goft *Goft) BasePath(group string) *Goft {
	goft.gg = baseGoftGroup(goft, group)

	return goft
}
```

在不使用 BasePath 的情况下， Mount 方法内会自建 `/` 路由作为 prefix。


## 使用

在 [main.go](/cmd/demo/main.go) 中，
1. 创建 BasePath Prefix， 名字为 `/demo`
2. 创建了 **2个** 路由组， `v1` 和 `v2`
    + 且 `v3` 是 `v2` 的子路由。

```go
func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()

	// 2. 设置 base Path
	g.BasePath("/demo")

	// 2. 注册多个路由组
	g.Mount("/v1", classes.NewIndex())

	{
		v2Router := g.Mount("/v2")
		v2Router.Mount("/v3", classes.NewIndex())
	}

	// 3. 启动 goft server
	g.Launch()
}
```

运行起来， 结果符合预期

```bash
cd cmd/demo/ && go run .
[GIN-debug] GET    /demo/v1/index            --> github.com/tangx-labs/gin-goft/cmd/demo/classes.handlerIndex (3 handlers)
[GIN-debug] GET    /demo/v2/v3/index         --> github.com/tangx-labs/gin-goft/cmd/demo/classes.handlerIndex (3 handlers)
[GIN-debug] Listening and serving HTTP on :8089
```

## 目录结构调整

1. 对 project 的名字进行了修改， 由 `goft` 改成 `demo` 。 
    1. 以避免和框架的 `goft` 产生语意冲突
    2. 该为 demo 能更好的说明这部分代码非框架 goft 的一部分。
2. 将 `classes` 移动到 `/cmd/demo/classes` 中。

这样目录结构就比较清晰了

1. `/goft` 是针对 gin 二次封装的框架代码
2. `/cmd/demo` 是使用 `goft` 框架的用例测试代码
3. `/docs` 是说明文档。

