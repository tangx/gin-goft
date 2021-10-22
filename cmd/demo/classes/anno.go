package classes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/goft"
	"github.com/tangx-labs/gin-goft/httpx"
)

type AnnoDemo struct {
	httpx.MethodGet

	Age *goft.Value `prefix:"user.age"`
}

func (demo *AnnoDemo) Path() string {
	return "/anno/demo"
}

func (demo *AnnoDemo) Handler(c *gin.Context) (interface{}, error) {

	msg := fmt.Sprintf("注解: %s", demo.Age.String())
	return msg, nil
}
