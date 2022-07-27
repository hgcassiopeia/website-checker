package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{Context: c}
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewGinContext(c))
	}
}

func (c *GinContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *GinContext) JSON(status int, v interface{}) {
	c.Context.JSON(status, v)
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080", //frontend
	}
	config.AllowHeaders = []string{
		"Origin",
	}
	r.Use(cors.New(config))

	return &MyRouter{r}
}
