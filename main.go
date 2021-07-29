package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
func GetLatestBlock(client ethclient.Client) *Block {
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())

	block, err := client.BlockByNumber(context.Background(), blockNumber)
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
			//To:       tx.To().String(),
		})
	}
	return _block
}
func printBlocks(blk []Block) {
	i := 0
	for _, bl := range blk {
		fmt.Println("------------------------------------Blok-----------------------------------", i)
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
func printLatestBlock(bl Block) {
	fmt.Println("---------------------------------Blok---------------------------------")
	fmt.Println("Blocknumber: ", bl.BlockNumber)
	fmt.Println("Timestamp: ", bl.Timestamp)
	fmt.Println("Difficulty: ", bl.Difficulty)
	fmt.Println("Hash :", bl.Hash)
	fmt.Println("Transactions count: ", bl.TransactionsCount)
	if bl.TransactionsCount != 0 {
		fmt.Println("Transactions of Block")
		for _, tx := range bl.Transactions {
			fmt.Println("       Hash: ", tx.Hash)
			fmt.Println("       Value: ", tx.Value)
			fmt.Println("       s: ", tx.Gas)
			fmt.Println("       Gas price: ", tx.GasPrice)
			fmt.Println("       Nonce: ", tx.Nonce)
			fmt.Println("       To: ", tx.To)
			fmt.Println("       Pending: ", tx.Pending)
		}
	}
}

func Transfer(client ethclient.Client, privKey string, to string, amount int64) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return "", err
	}

	value := big.NewInt(amount)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().String(), nil
}
func GetAddressBalance(client ethclient.Client, address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "0", err
	}

	return balance.String(), nil
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println()
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(client ethclient.Client) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Chose option(1 - Read last block, 2 - Make transfer, 3 - Account Balance)", reader)

	switch opt {
	case "1":

	case "2":

	case "3":

	default:

	}

}

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/d22481bed6d64ec39213f11d3050cb60")
	//client, err := ethclient.Dial("HTTP://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	//blocks := GetBlocks(*client)
	//block := GetLatestBlock(*client)
	//printLatestBlock(*block)
	fmt.Println(GetAddressBalance(*client, "0xE13EF9474558F84DC23D0fd4736b772aAdd0FD51"))
	//fmt.Println(Transfer(*client, "390ad60842a88911f019ad20782b32f07620107e2d07d0dc2b90cada46398829", "0xE13EF9474558F84DC23D0fd4736b772aAdd0FD51", 2322321312))

}
