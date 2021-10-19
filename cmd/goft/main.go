package main

import (
	"github.com/tangx-labs/gin-goft/classes"
	"github.com/tangx-labs/gin-goft/goft"
)

func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()

	g.BasePath("/demo")

	// 2. 注册路由
	g.Mount("/v1",
		classes.NewIndex(),
	)
	// 04.2. 注册多个路由组。
	g.Mount("/v2",
		classes.NewIndex(),
	).Mount("/v3",
		classes.NewIndex(),
	)

	// 3. 启动 goft server
	g.Launch()
}
