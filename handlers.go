package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckWallet(c *gin.Context) {
	pubkey := c.Query("pubkey")

	fmt.Println(pubkey)
	result, err := CheckWalletService(pubkey)

	if err != nil {
		fmt.Println(err)
		RespFail(c.Writer)
		return
	}

	fmt.Println(result)

	checkWallet := CheckWalletOutputDTO{
		IsExist: result,
	}

	RespSuccess(c.Writer, checkWallet)
}

func CreatedWallet(c *gin.Context) {
	fmt.Println("开始创建钱包")
	var createdWalletDTO CreatedWalletInputDTO
	if err := c.ShouldBindJSON(&createdWalletDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("userPublicKey:%s\n", createdWalletDTO.UserPubKey)

	_, err := ExecShellAndGetResult(CREATE_WALLET, []string{createdWalletDTO.UserPubKey})

	if err != nil {
		RespFail(c.Writer)
		return
	}

	addResult, err := SaveUserWallet(createdWalletDTO.UserPubKey, createdWalletDTO.UserPubKey)

	if err != nil {
		RespFail(c.Writer)
		return
	}

	// 创建响应对象
	RespSuccess(c.Writer, addResult)
}

func GetWalletBalance(c *gin.Context) {
	pubkey := c.Query("pubkey")

	fmt.Println(pubkey)

	shellResult, err := ExecShellAndGetResult(GET_BALANCE, []string{pubkey})
	if err != nil {
		RespFail(c.Writer)
	}

	// 创建响应对象
	RespSuccess(c.Writer, shellResult)
}

func GetReceiveAddress(c *gin.Context) {
	pubkey := c.Query("pubkey")

	fmt.Println(pubkey)

	shellResult, err := ExecShellAndGetResult(RECEIVE, []string{pubkey})
	if err != nil {
		RespFail(c.Writer)
		return
	}

	// 创建响应对象
	RespSuccess(c.Writer, shellResult)
}

func GetTransactions(c *gin.Context) {
	pubkey := c.Query("pubkey")

	fmt.Println(pubkey)

	shellResult, err := ExecShellAndGetResult(TRANSACTIONS, []string{pubkey})
	if err != nil {
		RespFail(c.Writer)
		return
	}

	// 创建响应对象
	RespSuccess(c.Writer, shellResult)
}

func SendRunes(c *gin.Context) {
	var sendRunesInputDTO SendRunesInputDTO
	if err := c.ShouldBindJSON(&sendRunesInputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shellResult, err := ExecShellAndGetResult(SEND, []string{sendRunesInputDTO.ReceiveAddress, sendRunesInputDTO.RunesName, fmt.Sprintf("%.2d", sendRunesInputDTO.RunesAmount), fmt.Sprintf("%.2f", sendRunesInputDTO.FeeRate)})

	if err != nil {
		RespFail(c.Writer)
		return
	}

	// 创建响应对象
	RespSuccess(c.Writer, shellResult)
}

func ImportWallet(c *gin.Context) {

}
