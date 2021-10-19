package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"hello": "gin-goft",
		})
	})

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
