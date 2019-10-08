package handlers

import (
	"bufio"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shrotavre/filetinder/internal/filetinder"
)

// GetTargets return gin handler to get stored targets
func GetTargets(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	c.JSON(http.StatusOK, targetStoreInst.List())
}

// GetTarget return gin handler to get single stored targets
func GetTarget(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	id, _ := strconv.Atoi(c.Param("id"))

	t := targetStoreInst.FindByID(id)
	if t != nil {
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}

// GetTargetFile return gin handler to get single stored target's file
func GetTargetFile(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	id, _ := strconv.Atoi(c.Param("id"))

	t := targetStoreInst.FindByID(id)
	if t != nil {
		f, err := t.GetFile()
		if err != nil {
			log.Panic(err)
			c.Status(http.StatusBadRequest)
			return
		}

		fr := bufio.NewReader(f)

		fstat, err := t.GetStats()
		if err != nil {
			log.Panic(err)
			c.Status(http.StatusBadRequest)
			return
		}

		extraHeaders := map[string]string{
			// "Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, fstat.Size, fstat.ContentType, fr, extraHeaders)

		defer f.Close()

		return
	}

	c.Status(http.StatusNotFound)
}

// GetTargetStat return gin handler to get single stored target's filestats
func GetTargetStat(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	id, _ := strconv.Atoi(c.Param("id"))

	t := targetStoreInst.FindByID(id)
	if t != nil {
		fstat, err := t.GetStats()
		if err != nil {
			log.Panic(err)
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, fstat)
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
		URL:  requestBody.URL,
		Tags: make([]string, 0),
	}

	// Add to store
	targetStoreInst := filetinder.TargetStoreInst
	svdt := targetStoreInst.Add(&t)

	c.JSON(http.StatusOK, svdt)
}

// DeleteTarget return gin handler to delete target
func DeleteTarget(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst
	id, _ := strconv.Atoi(c.Param("id"))

	t := targetStoreInst.FindByID(id)
	if t != nil {
		targetStoreInst.Del(t)
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}

// MarkTarget return gin handler to mark target
func MarkTarget(c *gin.Context) {
	targetStoreInst := filetinder.TargetStoreInst

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

	t := targetStoreInst.FindByID(id)
	if t != nil {
		if !t.HasTag(requestBody.Value) {
			t.Tags = append(t.Tags, requestBody.Value)
		}
		c.JSON(http.StatusOK, t)
		return
	}

	c.Status(http.StatusNotFound)
}
