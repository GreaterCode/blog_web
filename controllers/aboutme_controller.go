package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutMeGet(c *gin.Context) {
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "aboutme.html", gin.H{"IsLogin": islogin, "wechat": "微信：renwoxing", "qq": "QQ:123456",
		"tel": "Tel:1223344"})
}
