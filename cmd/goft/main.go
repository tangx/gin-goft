package main

import (
	"github.com/tangx-labs/gin-goft/classes"
	"github.com/tangx-labs/gin-goft/goft"
)

func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()

	// 2. 注册路由
	g.Mount(
		classes.NewIndex(),
	)

	// 3. 启动 goft server
	g.Launch()
}
