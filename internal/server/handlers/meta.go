package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/config"
)

type meta struct {
	Port int `json:"port"`
	PID  int `json:"pid"`
}

// GetMeta return gin handler to get server meta
func GetMeta(c *gin.Context) {
	appmeta := &meta{}
	appmeta.Port = config.DefaultPort
	appmeta.PID = os.Getpid()

	c.JSON(http.StatusOK, appmeta)
}
