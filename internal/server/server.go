package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/filetinder"
	"github.com/shrotavre/filetinder/internal/server/handlers"
)

func getCORSMiddleware() gin.HandlerFunc {
	CORSConfig := cors.DefaultConfig()
	CORSConfig.AllowOrigins = []string{"http://localhost:5000"}

	return cors.New(CORSConfig)
}

func setupRouter() *gin.Engine {
	if filetinder.Config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(getCORSMiddleware())

	apis := r.Group("/api")
	{
		// Targets
		apis.GET("targets", handlers.GetTargets)
		apis.POST("targets", handlers.AddTarget)
		apis.GET("targets/:id", handlers.GetTarget)
		apis.GET("targets/:id/file", handlers.GetTargetFile)
		apis.GET("targets/:id/fstats", handlers.GetTargetStat)
		apis.DELETE("targets/:id", handlers.DeleteTarget)
		apis.POST("targets/:id/mark", handlers.ToggleMarkTarget)

		// Funcs
		apis.POST("/funcs/stop-server", handlers.StopServerFunc)
		apis.POST("/funcs/delete-all", handlers.DeleteAllFunc)

		// Meta
		apis.GET("/meta", handlers.GetMeta)
	}

	r.StaticFS("/ui", http.Dir(filetinder.Config.UIPath))

	return r
}

// Start starts FileTinder main HTTP server
func Start() error {
	port := filetinder.Config.Port

	r := setupRouter()

	return r.Run(fmt.Sprintf(":%d", port))
}
