package main

import (
	"fmt"
	"log"
	"reflect"
)

func CheckWalletService(pubkey string) (bool, error) {
	sql := "select count(1) as total from tb_user_wallet_mapping where pubkey = ?"

	params := []interface{}{pubkey}
	return CheckExist(sql, params)
}

func SaveUserWallet(pubkey string, walletName string) (bool, error) {
	sql := "insert into tb_user_wallet_mapping(pubkey, wallet_name) values(?,?)"

	params := []interface{}{pubkey, walletName}

	result, err := SaveData(sql, params)
	if err != nil {
		log.Printf("Failed to add user wallet mapping: %v", err)
		return false, err
	}

	return result, nil
}

func SaveApiKey(callerId, apiKey string) (bool, error) {
	sql := "insert into tb_api_key (caller_id, api_key) values(?,?)"

	params := []interface{}{callerId, apiKey}

	result, err := SaveData(sql, params)
	if err != nil {
		log.Printf("Failed to add api-key: %v", err)
		return false, err
	}

	return result, nil
}

func (s *ApiKeyStore) LoadApiKeyToRedis() {
	sql := "select api_Key from tb_api_key"

	var apiKeyModel ApiKeyModel
	apiKeyList, _ := QueryDataList(sql, nil, apiKeyModel)

	// 使用反射来遍历 interface{}
	val := reflect.ValueOf(apiKeyList)

	// 确保 data 是一个切片
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			// 获取切片中的元素
			elem := val.Index(i).Interface()
			fmt.Println("Processing API Key:", elem)
		}
	} else {
		fmt.Println("Provided data is not a slice")
	}

	fmt.Println(apiKeyList)
}
