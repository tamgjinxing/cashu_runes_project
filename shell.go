package main

import (
	"fmt"
	"os/exec"
)

var scriptPath = "command.sh"
var testnetScriptPath = "testnet_command.sh"

const (
	SEND          string = "send"
	RECEIVE       string = "receive"
	TRANSACTIONS  string = "transactions"
	CREATE_WALLET string = "create"
	GET_BALANCE   string = "balance"
)

func ExecShellAndGetResult(handleType string, args []string) (interface{}, error) {
	// 假设有一个脚本文件 script.sh 需要传递参数
	fullArgs := append([]string{config.EnvInfo.SctiptName, handleType}, args...)

	// 定义要执行的命令，并传递参数
	cmd := exec.Command("sh", fullArgs...)
	// 获取命令的输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing script with arguments:", err)
		return nil, err
	}

	// 打印脚本输出
	fmt.Printf("Shell Output: %s\n", output)
	return output, nil
}
