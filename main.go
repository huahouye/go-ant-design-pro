package main

import (
	"embed"
	"go-ant-design-pro/internal/logger"
	"go-ant-design-pro/middleware"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//go:embed webui/dist
var webui embed.FS

func main() {

	logger.Info("Server starting at port 8080...")

	router := gin.Default()
	// logger, if you dont want to log static access, please move down these two lines
	logger.PluginInGin(router)
	// Serve frontend webui
	router.Use(static.Serve("/", middleware.NewEmbedFolder(webui, "webui/dist")))

	// Setup route group for the API v1
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	// subgroup
	customers := apiV1.Group("/customer")
	{
		customers.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello customer",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		logger.Infof("%s resource doesn't exists, redirect on page ", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// listen and serve on 0.0.0.0:8080
	router.Run()

}
