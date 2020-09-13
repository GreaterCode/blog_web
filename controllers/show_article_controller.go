package controllers

import (
	"blog_web/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowArticleGet(c *gin.Context) {
	// 获取session
	islogin := GetSession(c)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	art := models.QueryArticleWithId(id)
	c.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Content": art.Content})
}
