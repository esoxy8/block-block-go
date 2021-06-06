package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

/*
	Description:

	Combines the previous block's hash and the block's data to generate a
	new unique has for the block.

	Parameters: 	None
	Returns: 		None

*/
func (b *Block) HashBlock() {
	unhashed := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hashed := sha256.Sum256(unhashed)
	b.Hash = hashed[:]
}

/*
	Description: 		Creates a new block

	Parameters: 	1.	data (string)
					2. 	previous blocks hash ([]byte)

	Returns 		1.  Pointer to the newly created block

*/
func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{
		Data:     []byte(data),
		PrevHash: prevHash,
		Hash:     []byte{},
	}
	block.HashBlock()
	return &block
}
