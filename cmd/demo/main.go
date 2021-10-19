package main

import (
	"github.com/tangx-labs/gin-goft/cmd/demo/classes"
	"github.com/tangx-labs/gin-goft/cmd/demo/middlewares"
	"github.com/tangx-labs/gin-goft/goft"
)

func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()
	// g.Attach(&middlewares.User{})

	// 2. 设置 base Path
	g.BasePath("/demo")

	// 2. 注册多个路由组
	g.Mount("/v1", classes.NewIndex())

	{
		v2Router := g.Mount("/v2")
		// 子路由注册中间件
		v2Router.Attach(&middlewares.User{})

		v2Router.Mount("/v3", classes.NewIndex())

	}

	// 3. 启动 goft server
	g.Launch()
}
