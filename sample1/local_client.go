package main

import (
	"context"
	"fmt"
	"go-ethereum/crypto"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("/Users/ahirs/Library/Ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	//address := common.HexToAddress("c9e21e87ac7a95ec366c01e3888747cadc5b0ea5")

	account := common.HexToAddress("oxa26036c277c33f09d7aa94eabfea75a388b5b066")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("balance", balance)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time().Uint64())       // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())
	fmt.Println(block.Transactions())

	file1 := "UTC--2019-02-12T19-09-58.790935000Z--a26036c277c33f09d7aa94eabfea75a388b5b066"
	file2 := "UTC--2019-02-12T19-10-07.000403000Z--7fffeac89ac63fe54849385e729cc187323c0fa6"
	private, err := crypto.HexToECDSA("0xc692A6FD08c7B49a7e1B52b8F2D09B0D237AB851")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("private", private)

}
