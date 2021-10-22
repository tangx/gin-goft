package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	_wanted := "zhangsan"
	if user.Name != "zhangsan" {
		logrus.Warnf("user is %s, wanted %s", user.Name, _wanted)
	}

	return
}
