package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func InitRoute(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRouter := router.Group("/api")
	apiRouter.GET("/get-age", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"age": 18,
		})
	})
	apiRouter.GET("/get-name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "lts003",
		})
	})
	apiRouter.GET("/exit-system", func(c *gin.Context) {
		os.Exit(1)
	})
}
