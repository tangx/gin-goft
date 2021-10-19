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

func (gg *GoftGroup) Handle(class ClassController) {

	m := class.Method()
	p := class.Path()
	handler := class.Handler

	gg.RouterGroup.Handle(m, p, func(c *gin.Context) {
		err := ginbinder.ShouldBindRequest(c, class)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		v, err := handler()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, v)
	})

}
