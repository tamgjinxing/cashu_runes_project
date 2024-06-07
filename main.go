package main

import (
	"ecash_runes_project/storage"
	"log"
)

// func main() {
// 	fmt.Println("Welcome to ecash runes project!!!")

// 	SendRunes()
// 	ReceiveRunes()
// }

func main() {

	// 初始化配置文件
	InitConfig()

	//初始化并注册数据库管理器
	dbManager, err := storage.NewDBManager(config.DBInfo.DBUrl)
	if err != nil {
		log.Printf("Failed to initialize DB manager: %v", err)
	}
	defer dbManager.Close()
	storage.Register("DBManager", dbManager)

	// 启动http服务
	InitHttpServer()
}
