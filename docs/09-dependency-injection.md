# 依赖注入

不可避免的， 在开发过程中一定会操作 **数据库、缓存** 等其他中间件。


```go
type GetUserByID struct {
	httpx.MethodGet
	DBA *adaptors.GormAdaptor
}
func (user *GetUserByID) Handler(c *gin.Context) (interface{}, error) {
	um := &models.User{}
	// ... 省略
	user.DBA.DB.Where("user_id=?", params.UserID).First(um)
	return um, nil
}
```

那么如何相对优雅的获取 **操作句柄** 呢？ 


## 依赖注入

思想其实挺简单， 就是在 Goft 启动的时候， 在 **Bind** 环境， 在 **GoftGroup.adaptors** 中查找是否有 **元素** 匹配的 Class 控制器中的 **适配器(Adaptor)** 字段， 如果有则该元素绑定到 Class 控制器中。


### GoftGroup 中的 adaptor 容器

#### adaptor 数据类型选型

由于无法预先确认 adaptor 的类型， 所以只能使用 interface 表示。

```go
type GoftGroup struct {
	*gin.RouterGroup
	// 适配器， 类似数据库等
	adaptors []interface{}
}
type Goft struct {
	*gin.Engine
	rootGrp *GoftGroup

	onceWithAdaptors sync.Once
}
```

#### adaptor 数据共享注意

终端控制器（路由） 是被绑定到 GoftGroup 上的。 并且所有 GoftGroup 是一个 **链表**， 且链上的所有节点都使用相同的 adaptors 切片容器。 

```go
// baseGoftGroup 通过 Goft 返回一个根 GoftGroup
func baseGoftGroup(r *Goft, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: r.RouterGroup.Group(group),
		adaptors:    make([]interface{}, 0),
	}
}

// newGoftGroup 通过 GoftGroup 扩展新的 GoftGroup
func newGoftGroup(base *GoftGroup, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: base.Group(group),
		adaptors:    base.adaptors,
	}
}
```

为了保证链上的所有节点都具有 **相同的 adaptors**。  

1. `Goft.WithAdaptors()` **只能** 被执行一次， 这是因为 adaptor 使用的是切片扩容特定导致的。

```go
// WithAdaptors 注入适配器， 比如 *gorm.DB, *goredis.Redis
func (goft *Goft) WithAdaptors(adaptors ...interface{}) {
	goft.onceWithAdaptors.Do(
		func() {
			goft.rootGrp.adaptors = append(goft.rootGrp.adaptors, adaptors...)
		},
	)
}
```

2. `Goft.WithAdaptors()` **必须** 创建在任何 **子group** 之前。 这个在用户侧顺序， 无法通过代码约束。

```go
func main() {
	g := goft.Default()

	g.WithAdaptors(
		adaptors.NewGormAdaptor(),
	)

	demo := g.Mount("/demo")
}
```

### 实现注入

注入的过程就比较简单粗暴了。
1. 在 setAdaptor 中获取 class 中的所有字段 `for i := 0; i < rv.NumField(); i++`
2. 对字段进行有效性判断 `if fv.Kind() != reflect.Ptr || !fv.IsNil()`， 非 nil 字段表示用于已经对其赋值， 因此不再使用注入赋值。
3. 获取字段 reflect.Type， 通过 `getAdaptor()` 在 `adaptors` 中查找是否有 **相同类型** 的元素。
4. 如果有， 通过反射进行赋值

```go
		ft := fv.Type()
		if adp := gg.getAdaptor(ft); adp != nil {
			fv.Set(reflect.New(fv.Type().Elem()))
			fv.Elem().Set(reflect.ValueOf(adp).Elem())
		}
```

注意： 
1. reflect.New 返回的是对象指针， 所以要使用 `fv.Type().Elem()`。
2. 否则 `fv.Type()` 的类型可能是 `**adaptors.GormAdaptor`, 注意两个 `*` 号。 

上诉这种方法是替换 **指针后面的真实对象**， 当然也可以直接进行 **指针替换**

```go
		if adp := gg.getAdaptor(ft); adp != nil {
			// 此时字段是 nil ， 因此对字段 fv 先进行初始化
			// fv.Set(reflect.New(fv.Type().Elem()))
			// 反射赋值
			// fv.Elem().Set(reflect.ValueOf(adp).Elem())
			fv.Set(reflect.ValueOf(adp))
		}
```

完整代码如下

```go
// Bind 重载 GoftGroup 的 Bind 方法
func (gg *GoftGroup) Bind(class ClassController) IGoftRoutes {
	gg.setAdaptor(class)
	return gg
}

// setAdaptor 为 class 注入匹配的 adaptor
func (gg *GoftGroup) setAdaptor(class ClassController) {
	rv := reflect.ValueOf(class)
	rv = reflect.Indirect(rv)

	for i := 0; i < rv.NumField(); i++ {
		// 循环遍历所有字段
		fv := rv.Field(i)
		// 如果 fv 已有值， 或者 fv 不是指针，则跳过
		if fv.Kind() != reflect.Ptr || !fv.IsNil() {
			continue
		}

		ft := fv.Type()
		if adp := gg.getAdaptor(ft); adp != nil {
			fv.Set(reflect.New(fv.Type().Elem()))
			fv.Elem().Set(reflect.ValueOf(adp).Elem())
		}
	}
}

// getAdaptor 从 GoftGroup 中的 adaptors 中返回匹配类型的 Adaptor
func (gg *GoftGroup) getAdaptor(t reflect.Type) interface{} {
	for _, adaptor := range gg.adaptors {
		if t == reflect.TypeOf(adaptor) {
			return adaptor
		}
	}

	return nil
}
```

至此， 就可以向开篇一样开心的直接使用 **数据库具柄** 了。
