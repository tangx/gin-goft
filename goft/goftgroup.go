package goft

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type IGoftRouter interface {
	IGoftRoutes

	Mount(string, ...ClassController) *GoftGroup
}
type IGoftRoutes interface {
	Bind(ClassController) IGoftRoutes
	Attach(...Fairing) IGoftRoutes
}

var _ IGoftRouter = &GoftGroup{}

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
		grp.Bind(class)
	}

	return grp
}

// Attach 绑定/注册 中间件
func (gg *GoftGroup) Attach(fairs ...Fairing) IGoftRoutes {
	return gg.attach(fairs...)
}

func (gg *GoftGroup) attach(fairs ...Fairing) IGoftRoutes {
	for _, fair := range fairs {
		fair := fair

		// 创建一个临时中间件 handler
		handler := func(c *gin.Context) {

			// cc := c.Copy()
			// 这里不应该传入 cc 备份给 Middleware 处理。
			// 某些中间件可能就是需要修改 gin.Context 中的一些内容。
			// 如果要避免类似中间件读取 body， 而导致业务逻辑失效的话
			//    可以在 OnRequest 中自行使用 cc 副本
			// if err := fair.OnRequest(cc); err != nil {
			// 	// c.Abort()
			// 	c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			// 		"err": err.Error(),
			// 	})
			// 	return
			// }

			// 由于 goft 是一个框架， 不应该对任何已经放行的中间件做任何阻拦
			// 如果需要中断， 可以在业务实现的中间件本身中进行阻拦。
			_ = fair.OnRequest(c)
			c.Next()
		}

		// 使用 中间件
		gg.Use(handler)
	}

	return gg
}

// Bind 重载 GoftGroup 的 Bind 方法
func (gg *GoftGroup) Bind(class ClassController) IGoftRoutes {

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

	return gg
}
