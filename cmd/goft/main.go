package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx-labs/gin-goft/classes"
)

func main() {
	r := gin.Default()

	// 3. 向 gin engine 注册路由信息
	classes.NewIndex(r).Build()

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
