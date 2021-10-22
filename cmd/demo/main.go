package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx-labs/gin-goft/cmd/demo/adaptors"
	"github.com/tangx-labs/gin-goft/cmd/demo/annotations"
	"github.com/tangx-labs/gin-goft/cmd/demo/classes"
	"github.com/tangx-labs/gin-goft/cmd/demo/middlewares"
	"github.com/tangx-labs/gin-goft/goft"
)

func main() {

	// 1. 使用 goft 代替 gin
	g := goft.Default()

	// 注解答
	g.WithAnnotations(&annotations.Value{})

	// 适配器， 不如 数据库连接池 什么的
	g.WithAdaptors(adaptors.NewGormAdaptor())

	// 中间件
	g.Attach(&middlewares.User{})

	// 路由组
	{
		demo := g.Mount("/demo")
		// 2. 注册多个路由组
		demo.Mount("/v1",
			classes.NewIndex(),
			&classes.GetUserByID{},
			classes.NewAnnoDemo(),
		)

		v2 := demo.Mount("/v2")
		// 子路由注册中间件
		v2.Mount("/v3", classes.NewIndex())
	}

	// 3. 启动 goft server
	g.Launch()
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}
