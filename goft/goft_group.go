package goft

import "github.com/gin-gonic/gin"

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
func (gg *GoftGroup) Mount(group string, claess ...ClassController) *GoftGroup {
	grp := newGoftGroup(gg, group)

	for _, class := range claess {
		class.Build(grp)
	}

	return grp
}
