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

func hasTag(ts []string, s string) bool {
	for _, t := range ts {
		if t == s {
			return true
		}
	}
	return false
}

// DeleteAllFunc return gin handler to delete all marked files
func DeleteAllFunc(c *gin.Context) {
	ts := filetinder.TargetStore
	nts := make([]*filetinder.Target, 0)

	for _, t := range ts {
		removed := true

		if hasTag(t.Tags, "remove") {
			err := os.Remove(t.URL)
			if err != nil {
				log.Panic(err)
				removed = false
			}
		}

		if removed == false {
			nts = append(nts, t)
		}
	}

	ts = nts

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
