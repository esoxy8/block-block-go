package logging

import (
	"fmt"

	"github.com/esoxy8/block-block-go/blockchain"
)

func PrintBlock(b *blockchain.Block) {
	fmt.Printf("DATA: %s\n", b.Data)
	fmt.Printf("HASH: %x\n", b.Hash)
	fmt.Printf("PREVIOUS HASH: %x\n\n\n", b.PrevHash)
}
