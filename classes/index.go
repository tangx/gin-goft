package classes

import "github.com/gin-gonic/gin"

// 1. 创建一个 Index 业务控制器， 被在其中内嵌一个 gin engine

// Index
type Index struct {
	e *gin.Engine
}

func NewIndex(e *gin.Engine) *Index {
	return &Index{
		e: e,
	}
}

// Build 2. 构造器， 创建路由信息
func (index *Index) Build() {
	index.e.Handle("GET", "/", handlerIndex)
}

func handlerIndex(c *gin.Context) {
	c.JSON(200, map[string]string{
		"hello": "gin-goft",
	})
}
