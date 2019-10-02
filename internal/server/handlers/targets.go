package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	target struct {
		ID  int64  `json:"id"`
		URL string `json:"url" form:"url"`
		Tag string `json:"tag"`
	}

	targets []*target
)

var (
	targetStore       targets
	targetIDIncrement int64
)

func init() {
	targetStore = make([]*target, 0)
	targetIDIncrement = 1
}

func findTargetByID(id int64) *target {
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
	var t target

	err := c.BindJSON(&t)
	if err != nil {
		log.Panic(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Set default value
	t.ID = targetIDIncrement

	// Add to store
	targetStore = append(targetStore, &t)
	targetIDIncrement++

	c.JSON(http.StatusOK, targetStore)
}

// MarkTarget return gin handler to mark target
func MarkTarget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	t := findTargetByID(int64(id))
	if t != nil {
		t.Tag = "marked"

		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}
