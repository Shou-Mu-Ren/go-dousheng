package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var basePath = "/home/runner/app/public" 

func FileDownload(c *gin.Context) {
	filename := c.Query("filename")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(basePath + filename)
}
