package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
	rg *gin.RouterGroup
}

// Default 创建一个默认的 Engine
func Default() *Goft {
	return &Goft{
		Engine: gin.Default(),
	}
}

// NewWithEngine 使用自定义 gin engine 创建
func NewWithEngine(e *gin.Engine) *Goft {
	return &Goft{
		Engine: e,
	}
}

// Mount 挂载控制器
// 03.1. 关联控制器与 goft
// 03.2. 返回 *Goft 是为了方便链式调用
func (goft *Goft) Mount(group string, classes ...ClassController) *Goft {

	// 04.1. 注册路由组
	goft.rg = goft.Group(group)

	for _, class := range classes {
		// 03.3. 将 goft 传入到控制器中
		class.Build(goft)
	}

	return goft
}

// Handle 重载 gin.Engine 的 Handle 方法。
// 04.2. 这样子路由注册的时候， 就直接挂载到了 RouterGroup 上， 有了层级关系
func (goft *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	goft.rg.Handle(httpMethod, relativePath, handlers...)
}

// Launch 启动 gin-goft server。
// 这里由于重载问题， 不能将启动方法命名为 Run
func (goft *Goft) Launch() error {
	return goft.Run(":8089")
}
