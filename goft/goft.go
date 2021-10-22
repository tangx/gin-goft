package goft

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var _ IGoftRouter = &Goft{}

type Goft struct {
	*gin.Engine
	rootGrp *GoftGroup

	onceWithAdaptors sync.Once
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

	goft.initial()

	return goft
}

// initial 初始化 Goft
func (goft *Goft) initial() {
	if goft.rootGrp == nil {
		goft.rootGrp = baseGoftGroup(goft, "/")
	}
}

// Launch 启动 gin-goft server。
func (goft *Goft) Launch() error {
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	return goft.Run(addr)
	// return goft.Run(addrs...)
}

// Mount 挂载控制器
// 03.1. 关联控制器与 goft
// 03.2. 返回 *GoftGroup 是为了方便链式调用
func (goft *Goft) Mount(group string, classes ...ClassController) *GoftGroup {
	// 04.1. 注册路由组
	return goft.rootGrp.Mount(group, classes...)
}

func (goft *Goft) Attach(fairs ...Fairing) IGoftRoutes {
	return goft.rootGrp.Attach(fairs...)
}

func (goft *Goft) Bind(class ClassController) IGoftRoutes {
	return goft.rootGrp.Bind(class)
}

// WithAdaptors 注入适配器， 比如 *gorm.DB, *goredis.Redis
func (goft *Goft) WithAdaptors(adaptors ...interface{}) {
	goft.onceWithAdaptors.Do(
		func() {
			goft.rootGrp.adaptors = append(goft.rootGrp.adaptors, adaptors...)
		},
	)
}

func (goft *Goft) WithAnnotations(annos ...IAnnotation) *Goft {
	for _, anno := range annos {
		IAnnotationList = append(IAnnotationList, anno)
	}

	return goft
}
