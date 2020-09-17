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

		// 显示详情页
		router.GET("/show/:id", controllers.ShowArticleGet)

		//退出
		router.GET("/exit", controllers.ExitGet)

		v1 := router.Group("/article")
		{
			v1.GET("/add", controllers.AddArticleGet)
			v1.POST("/add", controllers.AddArticlePOST)

			// 显示文章内容
			v1.GET("/show/:id", controllers.ShowArticleGet)

			// 更新文章
			v1.GET("/update", controllers.UpdateArticleGet)
			v1.POST("/update", controllers.UpdateArticlePost)

			// 删除文章
			v1.DELETE("/delete", controllers.DeleteArticleGet)
		}

		// 标签
		router.GET("/tags", controllers.TagsGet)

		//相册
		router.GET("/album", controllers.AlbumGet)

		//文件上传
		router.POST("/upload", controllers.UploadPost)

		//关于我
		router.GET("/aboutme", controllers.AboutMeGet)
	}

	return router
}
