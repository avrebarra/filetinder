package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/filetinder"
)

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}

// DeleteAllFunc return gin handler to delete all marked files
func DeleteAllFunc(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	ts := targetStoreInst.List()

	for _, t := range ts {
		_, found := t.FindTag("remove")
		if found {
			err := os.Remove(t.URL)
			if err != nil {
				log.Panic(err)
			} else {
				targetStoreInst.Del(t)
			}
		}
	}

	c.Status(http.StatusOK)
}

// StopServerFunc return gin handler to stop server
func StopServerFunc(c *gin.Context) {
	c.Status(http.StatusOK)

	go func() {
		delaySecond(1)
		os.Exit(2)
	}()
}
