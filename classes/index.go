package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/goft"
)

// Index
// 删除 e *gin.Engine ， 删除强耦合关系
type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

// Build 控制器的构造器， 创建路由信息
// 1. 通过传参 解耦控制器和 gin server 的关系
// 2. 通过实现 ClassController 接口关联与 goft
func (index *Index) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", handlerIndex)
}

func handlerIndex(c *gin.Context) {
	c.JSON(200, map[string]string{
		"hello": "gin-goft",
	})
}
