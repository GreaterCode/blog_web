package controllers

import (
	"blog_web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TagsGet(c *gin.Context) {
	islogin := GetSession(c)

	tags := models.QueryArticleWithParam("tags")

	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": models.HandleTagsListData(tags), "IsLogin": islogin})

}
