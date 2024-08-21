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

	api := router.Group("/api")
	api.Use(ApiKeyMiddleware())
	{
		api.GET("/checkWallet", CheckWallet)
		api.POST("/createdWallet", CreatedWallet)
		api.GET("/getBalance", GetWalletBalance)
		api.GET("/getAddress", GetReceiveAddress)
		api.POST("/sendRunes", SendRunes)
		api.GET("/getTransactions", GetTransactions)
	}

	router.GET("/getApiKey", GetApiKeyHandler)

	port := "9901"

	log.Printf("Starting HTTP server at port:%s\n", port)
	router.Run(":" + port)
	log.Printf("Started HTTP server at port:%s\n", port)
}
