package goft

import "github.com/gin-gonic/gin"

type GoftGroup struct {
	*gin.RouterGroup
}

func baseGoftGroup(r *Goft, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: r.Group(group),
	}
}

func newGoftGroup(base *GoftGroup, group string) *GoftGroup {
	return &GoftGroup{
		RouterGroup: base.Group(group),
	}
}

func (gg *GoftGroup) Mount(group string, claess ...ClassController) *GoftGroup {
	grp := newGoftGroup(gg, group)

	for _, class := range claess {
		class.Build(grp)
	}

	return grp
}
