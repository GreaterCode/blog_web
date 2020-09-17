package controllers

import (
	"blog_web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlbumGet(c *gin.Context) {
	islogin := GetSession(c)
	albums, _ := models.FindAllAlbums()
	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": islogin, "Album": albums})
}
