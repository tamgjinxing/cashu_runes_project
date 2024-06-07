package main

import "log"

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
