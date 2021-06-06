package main

import (
	"github.com/esoxy8/block-block-go/blockchain"
	logging "github.com/esoxy8/block-block-go/services/logger"
)

func main() {
	// initialize the blockchain
	chainPtr := blockchain.InitBlockchain()
	chainPtr.AddBlock("1st block")
	chainPtr.AddBlock("2nd block")

	for _, block := range chainPtr.Blocks {
		logging.PrintBlock(block)
	}
}
