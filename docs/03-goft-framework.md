# 初始化 goft 框架

创建 [/goft](/goft) 目录， 将所有 goft 框架相关的东西都放在这里。


## goft

创建 goft 对象, **匿名嵌套**  `gin.Engine`， 这样 Goft 对象就可以直接调用 gin.Engine 的所有方法了。

```go
type Goft struct {
	*gin.Engine
}
```

创建 Goft 对象提供了 **2种方式**。  Default 和 自定义。 


```go
// Default 创建一个默认的 Engine
func Default() *Goft {
	return &Goft{
		Engine: gin.Default(),
	}
}
```
不管哪种方式， 在注册的时候， 在初始化 Goft 的时候， 都通过 `Engine: gin.Default()` 指定字段对象。  因此 **匿名嵌套字段 / 命名字段** 使用同样的方法方式初始化。这个之前困扰了我一段时间。

## 解耦 goft 与控制器

之前在控制器中， 是将 `*gin.Engine` 作为控制器的嵌套字段的， 显然耦合太紧，不合适。 因此要想办法解决。

[/classes/index.go](/classes/index.go) ， 在 Index 中， 删除了 `*gin.Engine` 的嵌套， 取而代之的是将 **goft.Goft** 在 `Build()` 方法中从外部传入， 从而达到解耦的目的。

```go
// Index
// 删除 e *gin.Engine ， 删除强耦合关系
type Index struct {
	// e *gin.Engine
}

// Build 控制器的构造器， 创建路由信息
// 1. 通过传参 解耦控制器和 gin server 的关系
// 2. 通过实现 ClassController 接口关联与 goft
func (index *Index) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", handlerIndex)
}
```

为了管理很多不同的控制器， goft 实现了 [class_controller.go](/goft/class_controller.go) **ClassController** 接口

```go
type ClassController interface {
	Build(goft *Goft)
}
```


只要满足该接口的对象，都能通过 **Mount** 方法实现路由注册。

```go
// Mount 挂载控制器
// 1. 关联控制器与 goft
// 2. 返回 *Goft 是为了方便链式调用
func (goft *Goft) Mount(classes ...ClassController) *Goft {
	for _, class := range classes {

		// 将 goft 传入到控制器中
		class.Build(goft)
	}

	return goft
}
```

## main.go 修改

现在就可以使用 goft 代替 gin 了。

```go
func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()

	// 2. 注册路由
	g.Mount(
		classes.NewIndex(),
	)

	// 3. 启动 goft server
	g.Launch()
}
```

