package main

import (
	"blog_web/routers"
)

func main() {
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "./static")
	router.Run(":8080")
}
