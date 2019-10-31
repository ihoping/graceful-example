package main

import (
	"github.com/gin-gonic/gin"
	"graceful-example/graceful"
	"graceful-example/router"
)

func main() {
	engine := gin.Default()
	router.InitRoute(&engine.RouterGroup)
	graceful.ExampleHttp(engine)
}
