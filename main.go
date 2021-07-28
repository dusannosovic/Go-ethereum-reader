package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlocks(client ethclient.Client) []Block {
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	blockList := []Block{}
	var i int64
	for i = 0; i <= blockNumber.Int64(); i++ {

		block, err := client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		_block := &Block{
			BlockNumber:       block.Number().Int64(),
			Timestamp:         block.Time(),
			Difficulty:        block.Difficulty().Uint64(),
			Hash:              block.Hash().String(),
			TransactionsCount: len(block.Transactions()),
			Transactions:      []Transaction{},
		}
		for _, tx := range block.Transactions() {
			_block.Transactions = append(_block.Transactions, Transaction{
				Hash:     tx.Hash().String(),
				Value:    tx.Value().String(),
				Gas:      tx.Gas(),
				GasPrice: tx.GasPrice().Uint64(),
				Nonce:    tx.Nonce(),
				To:       tx.To().String(),
			})
		}
		blockList = append(blockList, *_block)
	}
	return blockList
}

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	blocks := GetBlocks(*client)
	i := 0
	for _, bl := range blocks {
		fmt.Println("Blok: ", i)
		fmt.Println("Blocknumber: ", bl.BlockNumber)
		fmt.Println("Timestamp: ", bl.Timestamp)
		fmt.Println("Difficulty: ", bl.Difficulty)
		fmt.Println("Hash :", bl.Hash)
		fmt.Println("Transactions count: ", bl.TransactionsCount)
		fmt.Println("Transactions of Block[", i, "]")
		for _, tx := range bl.Transactions {
			fmt.Println("       Hash: ", tx.Hash)
			fmt.Println("       Value: ", tx.Value)
			fmt.Println("       s: ", tx.Gas)
			fmt.Println("       Gas price: ", tx.GasPrice)
			fmt.Println("       Nonce: ", tx.Nonce)
			fmt.Println("       To: ", tx.To)
			fmt.Println("       Pending: ", tx.Pending)
		}
		i++
		fmt.Println()
	}
}
