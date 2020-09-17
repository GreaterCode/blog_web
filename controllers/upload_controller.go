package controllers

import (
	"blog_web/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadPost(c *gin.Context) {
	fileHeader, err := c.FormFile("upload")
	if err != nil {
		responseErr(c, err)
	}

	now := time.Now()
	fileType := "other"
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == "*.png" || fileExt == ".gif" || fileExt == ".jpg" {
		fileType = "img"
	}

	// 文件夹路径
	fileDir := fmt.Sprintf("/static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.Mkdir(fileDir, os.ModePerm)
	if err != nil {
		responseErr(c, err)
		return
	}

	timestamp := time.Now().Unix()
	fileName := fmt.Sprintln("%d-%s", timestamp, fileHeader.Filename)
	filepathStr := filepath.Join(fileDir, fileName)

	c.SaveUploadedFile(fileHeader, filepathStr)

	if fileType == "img" {
		album := models.Album{0, filepathStr, fileName, 0, timestamp}
		models.InsertAlbum(album)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})
}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})

}
