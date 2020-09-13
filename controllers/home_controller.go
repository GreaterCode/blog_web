package controllers

import (
	"blog_web/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeGet(c *gin.Context) {
	// 获取session，判断用户是否登录
	islogin := GetSession(c)

	tag := c.Query("tag")

	page, _ := strconv.Atoi(c.Query("page"))

	var (
		artList    []models.Article
		hasFoolter bool
	)

	if len(tag) > 0 {
		artList, _ = models.QueryArticlesWithTag(tag)
		hasFoolter = false
	} else {
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticleWithPage(page)
		hasFoolter = true
	}

	fmt.Println("page---1->", page)
	fmt.Printf("%d", len(artList))
	for _, x := range artList {
		fmt.Println(" v:%v", x)
	}
	homeFooterPageCode := models.ConfigHomeFooterPageCode(page)

	html := models.MakeHomeBlocks(artList, islogin)
	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": islogin, "Content": html, "hasFooter": hasFoolter, "PageCode": homeFooterPageCode})
}
