package classes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/cmd/demo/annotations"
	"github.com/tangx-labs/gin-goft/httpx"
)

type AnnoDemo struct {
	httpx.MethodGet

	Age *annotations.Value `prefix:"user.age"`
}

func NewAnnoDemo() *AnnoDemo {
	return &AnnoDemo{}
}

func (demo *AnnoDemo) Path() string {
	return "/anno/demo"
}

func (demo *AnnoDemo) Handler(c *gin.Context) (interface{}, error) {

	msg := fmt.Sprintf("注解: %s", demo.Age.String())
	return msg, nil
}
