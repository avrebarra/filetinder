package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}

// DeleteAllFunc return gin handler to delete all marked files
func DeleteAllFunc(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// StopServerFunc return gin handler to stop server
func StopServerFunc(c *gin.Context) {
	c.Status(http.StatusOK)

	go func() {
		delaySecond(1)
		os.Exit(2)
	}()
}
