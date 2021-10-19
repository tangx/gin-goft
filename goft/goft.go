package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
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
// 1. 关联控制器与 goft
// 2. 返回 *Goft 是为了方便链式调用
func (goft *Goft) Mount(classes ...ClassController) *Goft {
	for _, class := range classes {

		// 将 goft 传入到控制器中
		class.Build(goft)
	}

	return goft
}

// Launch 启动 gin-goft server。
// 这里由于重载问题， 不能将启动方法命名为 Run
func (goft *Goft) Launch() error {
	return goft.Run(":8089")
}
