package goft

import "github.com/gin-gonic/gin"

type ClassController interface {
	Method() string
	Path() string
	Handler(c *gin.Context) (interface{}, error)
}

type HandlerFunc = func() (interface{}, error)
