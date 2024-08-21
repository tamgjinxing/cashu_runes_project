package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config 结构体用于存储配置信息
type Config struct {
	Redis struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"Redis"`

	DBInfo struct {
		Url      string `json:"url"`
		Port     int    `json:"port"`
		DBName   string `json:"dbName"`
		UserName string `json:"userName"`
		Password string `json:"password"`
		DBUrl    string `json:"dbUrl"`
	} `json:"DBInfo"`

	EnvInfo struct {
		Env        string `json:"env"`
		SctiptName string `json:"sctiptName"`
	} `json:"EnvInfo"`
}

var config Config // 全局变量用于存储配置信息

// readConfig 函数用于读取配置文件并将其内容存储在全局变量中
func ReadConfig(filename string) error {
	log.Printf("加载配置文件:%s\n", filename)
	// 读取配置文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// 解析配置文件内容到 Config 结构体
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}
