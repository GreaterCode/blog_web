package main

import (
	"blog_web/database"
	"blog_web/routers"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "./static")
	router.Run(":8080")
}
