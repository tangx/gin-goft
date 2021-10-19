package goft

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// fairing 中间件封装， 可以理解 fairing 是中间件的控制器
// fairing 只是提供中间件应该使用的业务逻辑， 但他并不是中间件的 handler
type Fairing interface {
	// 这里使用 *gin.Context 作为参数， 为了方便以后在中间件处理的时候获取请求体中的信息
	OnRequest(c *gin.Context) error
}

func attachFairings(iroute gin.IRoutes, fairs ...Fairing) {
	for _, fair := range fairs {
		fair := fair

		// 创建一个临时中间件 handler
		handler := func(c *gin.Context) {

			// cc := c.Copy()
			// 这里不应该传入 cc 备份给 Middleware 处理。
			// 某些中间件可能就是需要修改 gin.Context 中的一些内容。
			// 如果要避免类似中间件读取 body， 而导致业务逻辑失效的话
			//    可以在 OnRequest 中自行使用 cc 副本
			if err := fair.OnRequest(c); err != nil {
				// c.Abort()
				c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
					"err": err.Error(),
				})
				return
			}
			c.Next()
		}

		// 使用 中间件
		iroute.Use(handler)
	}
}
