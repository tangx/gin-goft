package goft

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type GoftGroup struct {
	*gin.RouterGroup
}

// baseGoftGroup 通过 Goft 返回一个根 GoftGroup
func baseGoftGroup(r *Goft, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: r.Group(group),
	}
}

// newGoftGroup 通过 GoftGroup 扩展新的 GoftGroup
func newGoftGroup(base *GoftGroup, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: base.Group(group),
	}
}

// Mount 在 GoftGroup 上绑定/注册 控制器
func (gg *GoftGroup) Mount(group string, classes ...ClassController) *GoftGroup {
	grp := newGoftGroup(gg, group)

	for _, class := range classes {
		grp.Handle(class)
	}

	return grp
}

// Attach 绑定/注册 中间件
func (gg *GoftGroup) Attach(fairs ...Fairing) {
	attachFairings(gg, fairs...)
}

// Handle 重载 GoftGroup 的 Handle 方法
func (gg *GoftGroup) Handle(class ClassController) {

	m := class.Method()
	p := class.Path()
	handler := class.Handler

	// 将业务逻辑封装成为 gin.HandlerFunc
	handlerFunc := func(c *gin.Context) {
		// 绑定参数到对象中
		err := ginbinder.ShouldBindRequest(c, class)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// 执行业务逻辑，获取返回值
		v, err := handler()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// 以 JSON 格式返回信息
		c.JSON(http.StatusOK, v)
	}

	// 调用 gin RouterGroup 的 Handle 方法注册路由
	gg.RouterGroup.Handle(m, p, handlerFunc)
}
