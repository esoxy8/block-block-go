package blockchain

type BlockChain struct {
	Blocks []*Block
}

/*
	Description:

	Takes some data, creates a new block, and adds it to the blockchain.

	Parameters: 	1. 	data (string)
	Returns: 		None
*/
func (chain *BlockChain) AddBlock(data string) {
	previousBlocksHash := chain.Blocks[len(chain.Blocks)-1].Hash
	chain.Blocks = append(chain.Blocks, CreateBlock(data, previousBlocksHash))
}

func InitBlockchain() *BlockChain {
	// create the genesis block
	genesisPtr := CreateBlock("Genesis Block", []byte{})
	chain := BlockChain{Blocks: []*Block{genesisPtr}}
	return &chain
}
