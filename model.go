package main

type SendRunesInputDTO struct {
	ReceiveAddress string  `form:"receiveAddress" json:"receiveAddress"`
	RunesName      string  `form:"runesName" json:"runesName"`
	RunesAmount    uint64  `form:"runesAmount" json:"runesAmount"`
	FeeRate        float32 `form:"feeRate" json:"feeRate"`
}

type CreatedWalletInputDTO struct {
	UserPubKey string `form:"userPubKey" json:"pubkey"`
}

type CheckWalletOutputDTO struct {
	IsExist bool `json:"isExist"`
}
