package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGET(c *gin.Context) {
	// 返回html页面
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}
