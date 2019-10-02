package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/server/handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	apis := r.Group("/api")
	{
		// Targets
		apis.GET("targets", handlers.GetTargets)
		apis.POST("targets", handlers.AddTarget)
		apis.GET("targets/:id", handlers.GetTarget)
		apis.POST("targets/:id/mark", handlers.MarkTarget)

		// Funcs
		apis.POST("/funcs/delete-all", nil)

		// Meta
		apis.GET("/meta", handlers.GetMeta)
	}

	return r
}

// Start starts FileTinder main HTTP server
func Start() error {
	port := config.DefaultPort

	r := setupRouter()

	return r.Run(fmt.Sprintf(":%d", port))
}
