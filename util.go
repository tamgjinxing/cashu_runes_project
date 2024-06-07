package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitConfig() {
	// 读取配置文件
	err := ReadConfig("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func InitHttpServer() {
	// 1.创建路由
	router := gin.Default()

	project := router.Group("/api")
	{
		project.GET("/checkWallet", CheckWallet)
		project.GET("/createdWallet", CreatedWallet)
		project.GET("/getBalance", GetWalletBalance)
		project.GET("/getAddress", GetReceiveAddress)
		project.POST("/sendRunes", SendRunes)
		project.POST("/createdWallet", CreatedWallet)
	}

	port := "9901"

	// 3.监听端口，默认在8080
	log.Printf("Starting HTTP server at port:%s\n", port)
	router.Run(":" + port)
	log.Printf("Started HTTP server at port:%s\n", port)
}
