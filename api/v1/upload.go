package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/pkg/e"
	"goblog/utils"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := utils.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": e.GetMsg(code),
		"url":     url,
	})
}
