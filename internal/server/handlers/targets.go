package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/filetinder"
)

var (
	targetStore       filetinder.TargetsCollection
	targetIDIncrement int64
)

func init() {
	targetStore = make([]*filetinder.Target, 0)
	targetIDIncrement = 1
}

func findTargetByID(id int64) *filetinder.Target {
	for _, value := range targetStore {
		if value.ID == int64(id) {
			return value
		}
	}

	return nil
}

// GetTargets return gin handler to get stored targets
func GetTargets(c *gin.Context) {
	c.JSON(http.StatusOK, targetStore)
}

// GetTarget return gin handler to get single stored targets
func GetTarget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	t := findTargetByID(int64(id))
	if t != nil {
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}

// AddTarget return gin handler to add target
func AddTarget(c *gin.Context) {
	var t filetinder.Target

	err := c.BindJSON(&t)
	if err != nil {
		log.Panic(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Set default value
	t.ID = targetIDIncrement
	t.Tags = make([]string, 0)

	// Add to store
	targetStore = append(targetStore, &t)
	targetIDIncrement++

	c.JSON(http.StatusOK, t)
}

// MarkTarget return gin handler to mark target
func MarkTarget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	t := findTargetByID(int64(id))
	if t != nil {
		t.Tags = append(t.Tags, "marked")
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}
