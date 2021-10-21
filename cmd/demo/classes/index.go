package classes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/httpx"
)

// Index
type Index struct {
	httpx.MethodGet
	Name string `query:"name"`
}

func NewIndex() *Index {
	return &Index{}
}

func (index *Index) Path() string {
	return "/index"
}

// wanted
func (index *Index) Handler(c *gin.Context) (interface{}, error) {
	index.Name = c.Query("name")

	if index.Name != "wangwu" {
		return nil, fmt.Errorf("invalid user: %s", index.Name)
	}

	return "hello, gin-goft", nil
}
