package main

import (
	"file-store/lib"
	"file-store/model/mysql"
	"file-store/router"
	"log"
)

func main() {
	serverConfig := lib.LoadServerConfig() //初始化配置
	mysql.InitDB(serverConfig)             //初始化并链接mysql
	defer mysql.DB.Close()
	r := router.SetupRoute()
	r.LoadHTMLGlob("view/*")
	r.Static("/static", "./static")
	if err := r.Run(":80"); err != nil {
		log.Fatal("服务器启动失败...")
	}
}
