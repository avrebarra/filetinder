package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAllFunc return gin handler to delete all marked files
func DeleteAllFunc(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
