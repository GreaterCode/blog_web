package controllers

import (
	"blog_web/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
当访问/add路径的时候回触发AddArticleGet方法
响应的页面是通过HTML
*/
func AddArticleGet(c *gin.Context) {
	// 获取session
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin})
}

func AddArticlePOST(c *gin.Context) {
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	art := models.Article{0, title, tags, short, content,
		"go从入门到精通", time.Now().Unix()}

	_, err := models.AddArticle(art)

	// 返回数据给浏览器
	res := gin.H{}
	if err == nil {
		res = gin.H{"code": 1, "message": "ok"}
	} else {
		res = gin.H{"code": 0, "message": "error"}
	}

	c.JSON(http.StatusOK, res)
}
