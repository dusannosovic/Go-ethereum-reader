package main

type Block struct {
	BlockNumber       int64
	Timestamp         uint64
	Difficulty        uint64
	Hash              string
	TransactionsCount int
	Transactions      []Transaction
}

type Transaction struct {
	Hash     string
	Value    string
	Gas      uint64
	GasPrice uint64
	Nonce    uint64
	To       string
	Pending  bool
}
