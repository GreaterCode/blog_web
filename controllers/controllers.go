package controllers

import (
	"blog_web/models"
	"blog_web/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterGET(c *gin.Context) {
	// 返回html页面
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

// 处理注册
func RegisterPost(c *gin.Context) {
	fmt.Println("==================")
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已经存在"})
		return
	}

	// 注册用户名和密码
	// 对存储密码进行md5加密
	password = utils.MD5(password)
	fmt.Println("md5加密后：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
	}

}

func LoginGet(c *gin.Context) {
	// 返回html
	fmt.Println("登录")
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, "password:", password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		/*
		   设置了session后,将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
		   因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		session := sessions.Default(c)
		session.Set("loginuser", username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}
