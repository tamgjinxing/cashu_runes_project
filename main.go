package main

import (
	"ecash_runes_project/storage"
	"fmt"
	"log"
)

// func main() {
// 	fmt.Println("Welcome to ecash runes project!!!")

// 	SendRunes()
// 	ReceiveRunes()
// }xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

func main() {

	// 初始化配置文件
	InitConfig()

	fmt.Println(config.DBInfo.Url)

	//初始化并注册数据库管理器
	dbManager, err := storage.NewDBManager(config.DBInfo.DBUrl)
	if err != nil {
		log.Printf("Failed to initialize DB manager: %v", err)
	}
	defer dbManager.Close()
	storage.Register("DBManager", dbManager)

	// start  http server
	InitHttpServer()
}
