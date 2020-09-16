package controllers

import (
	"blog_web/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteArticleGet(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
