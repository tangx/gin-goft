package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
	gg *GoftGroup
}

// Default 创建一个默认的 Engine
func Default() *Goft {
	r := gin.Default()
	return NewWithEngine(r)
}

// NewWithEngine 使用自定义 gin engine 创建
func NewWithEngine(e *gin.Engine) *Goft {
	goft := &Goft{
		Engine: e,
	}

	return goft
}

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

// BasePath 设置 Goft 的根路由
func (goft *Goft) BasePath(group string) *Goft {
	goft.gg = baseGoftGroup(goft, group)

	return goft
}

// Launch 启动 gin-goft server。
// 这里由于重载问题， 不能将启动方法命名为 Run
func (goft *Goft) Launch() error {
	return goft.Run(":8089")
}
