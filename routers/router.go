package routers

import (
	"blog_web/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//加载viwes
	router.LoadHTMLGlob("views/*")

	//注册
	router.GET("/register", controllers.RegisterGET)

	return router
}
