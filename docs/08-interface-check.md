# 接口实现检查

为了保证路由树的正常扩展， 在 Goft 设计中嵌入了两个字段。

```go
type Goft struct {
	*gin.Engine
	rootGrp *GoftGroup
}
```

因为 `Goft` 也需有实现 `Mount, Attach` 等方法， 而这些方法都需要落在 `rootGrp *GoftGroup` 字段上。

```go
func (goft *Goft) Mount(group string, classes ...ClassController) *GoftGroup {
	// 04.1. 注册路由组
	return goft.rootGrp.Mount(group, classes...)
}

// ... 省略
```

为了保障 Goft 和 GoftGroup 实现的方法一致， 就需要对这两个类型进行 **接口实现检查**。 
而在 go 中， 类与接口并无直接关系， **一个类** 只要实现了 **某个接口** 的所有方法， 那这个类就是该接口的实现。 为了保障在编码阶段就能检查， go 提供了一个特殊的语法糖。

语法如下

```go
var _ interfaceName = StructInstance
```

使用如下

```go
type IGoftRouter interface {
	IGoftRoutes

	Mount(string, ...ClassController) *GoftGroup
}
type IGoftRoutes interface {
	Bind(ClassController) IGoftRoutes
	Attach(...Fairing) IGoftRoutes
}

var _ IGoftRouter = &GoftGroup{}
var _ IGoftRouter = &Goft{}
```

在 gin 中也有同样的检查。



