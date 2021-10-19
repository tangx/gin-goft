package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `query:"name"`
}

func NewUser() *User {
	return &User{}
}

// OnRequest 实现 Fairing 接口
// 这里是否应该使用 指针方法 呢？
//    即 `func (user User) OnRequest(c *gin.Context)`
func (user User) OnRequest(c *gin.Context) (err error) {

	user.Name = c.Query("name")
	if user.Name != "zhangsan" {
		err = fmt.Errorf("非法用户: %s", user.Name)

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	return
}
