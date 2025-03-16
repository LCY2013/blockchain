package main

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	// 设置 RPC 客户端连接配置
	connCfg := &rpcclient.ConnConfig{
		Host:         "192.168.3.131:18443", // 默认 regtest 端口
		User:         "fufeng",              // 替换为你的 RPC 用户名
		Pass:         "123456",              // 替换为你的 RPC 密码
		HTTPPostMode: true,                  // 使用 HTTP POST 模式
		DisableTLS:   true,                  // 禁用 TLS（仅用于本地测试）
	}

	// 创建 RPC 客户端
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatalf("Failed to create RPC client: %v", err)
	}
	defer client.Shutdown()

	// 1. 创建新地址
	address, err := client.GetNewAddress("")
	if err != nil {
		log.Fatalf("Failed to create new address: %v", err)
	}
	fmt.Printf("New Address: %s\n", address.String())

	// 2. 转账
	amount := 1.00 // 转账金额
	newAmount, err := btcutil.NewAmount(amount)
	if err != nil {
		log.Fatalf("Failed to convert amount: %v", err)
	}
	txID, err := client.SendToAddress(address, newAmount)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Printf("Transaction ID: %s\n", txID)

	// 查询UTXO
	utxos, err := client.ListUnspent()
	if err != nil {
		log.Fatalf("Failed to list unspent UTXOs: %v", err)
	}
	for _, utxo := range utxos {
		utxoStr, _ := json.Marshal(utxo)
		fmt.Printf("UTXO: %s\n", string(utxoStr))
	}

	// 3. 生成区块确认交易
	//blockHashes, err := client.Generate(1)
	blockHashes, err := client.GenerateToAddress(1, address, nil)
	if err != nil {
		log.Fatalf("Failed to generate block: %v", err)
	}
	fmt.Printf("Generated Block Hash: %s\n", blockHashes[0])

	utxos, err = client.ListUnspent()
	if err != nil {
		log.Fatalf("Failed to list unspent UTXOs: %v", err)
	}
	for _, utxo := range utxos {
		utxoStr, _ := json.Marshal(utxo)
		fmt.Printf("UTXO: %s\n", string(utxoStr))
	}

	// 4. 查询余额
	balance, err := client.GetBalance("*")
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("Current Balance: %f BTC\n", balance.ToBTC())
}
