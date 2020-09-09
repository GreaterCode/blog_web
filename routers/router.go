package routers

import (
	"blog_web/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//加载viwes
	router.LoadHTMLGlob("views/*")

	// 设置session middleware
	store := cookie.NewStore([]byte(("loginuser")))
	router.Use(sessions.Sessions("mysession", store))
	{
		//注册
		router.GET("/register", controllers.RegisterGET)
		router.POST("/register", controllers.RegisterPost)

		// 登录
		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		//首页
		router.GET("/", controllers.HomeGet)
	}

	return router
}
