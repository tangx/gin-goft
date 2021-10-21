package main

import (
	"github.com/tangx-labs/gin-goft/cmd/demo/classes"
	"github.com/tangx-labs/gin-goft/goft"
)

func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()
	// g.Attach(&middlewares.User{})

	demo := g.Mount("/demo")

	// 2. 注册多个路由组
	demo.Mount("/v1",
		classes.NewIndex(),
		&classes.GetUserByID{},
	)

	{
		v2 := demo.Mount("/v2")
		// 子路由注册中间件
		v2.Mount("/v3", classes.NewIndex())

	}

	// 3. 启动 goft server
	g.Launch(":8089")
}
