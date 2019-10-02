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

func findTargetByID(id int64) (index int, t *filetinder.Target) {
	for i, value := range targetStore {
		if value.ID == int64(id) {
			return i, value
		}
	}

	return -1, nil
}

// GetTargets return gin handler to get stored targets
func GetTargets(c *gin.Context) {
	c.JSON(http.StatusOK, targetStore)
}

// GetTarget return gin handler to get single stored targets
func GetTarget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, t := findTargetByID(int64(id))
	if t != nil {
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}

// AddTarget return gin handler to add target
func AddTarget(c *gin.Context) {
	var requestBody struct {
		URL string `json:"url"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Panic(err)
		c.Status(http.StatusBadRequest)
		return
	}

	t := filetinder.Target{
		ID:   targetIDIncrement,
		URL:  requestBody.URL,
		Tags: make([]string, 0),
	}

	// Add to store
	targetStore = append(targetStore, &t)
	targetIDIncrement++

	c.JSON(http.StatusOK, t)
}

// DeleteTarget return gin handler to delete target
func DeleteTarget(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	i, t := findTargetByID(int64(id))
	if t != nil {
		targetStore = append(targetStore[:i], targetStore[i+1:]...)
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}

// MarkTarget return gin handler to mark target
func MarkTarget(c *gin.Context) {
	var requestBody struct {
		Value string `json:"value"`
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Panic(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// set default value
	if requestBody.Value == "" {
		requestBody.Value = "marked"
	}

	_, t := findTargetByID(int64(id))
	if t != nil {
		t.Tags = append(t.Tags, requestBody.Value)
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}
