package controllers

import (
	"blog_web/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateArticleGet(c *gin.Context) {
	// 获取session
	islogin := GetSession(c)
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id----> %d", id)

	// 获取id 对应文章
	art := models.QueryArticleWithId(id)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Tags": art.Tags, "Short": art.Short, "Content": art.Content, "Id": art.Id})

}

// 修改文章
func UpdateArticlePost(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	// 获取浏览器传输的数据
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")

	// 实例化model,修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err = models.UpdateArticle(art)

	// 返回数据结构给浏览器
	res := gin.H{}
	if err != nil {
		res = gin.H{"code": 1, "message": "更新成功"}
	} else {
		res = gin.H{"code": 0, "message": "更新失败"}
	}

	c.JSON(http.StatusOK, res)

}
