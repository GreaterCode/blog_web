package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(c *gin.Context) {

	//清除用户session
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()

	fmt.Println("delete session:", session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently, "/")
}
