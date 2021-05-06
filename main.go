package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)


type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
type BlockChain struct{
	blocks []*Block
}

type ProofOfWork struct {
	Block *Block
	Target *big.Int
 }
 
func (b *Block) putHash() {
	info := append(b.Data, b.PrevHash...)
	// This will join our previous block's relevant info with the new blocks
	hash := sha256.Sum256(info)
	//This performs the actual hashing algorithm
	b.Hash = hash[:] 
	//If this ^ doesn't make sense, you can look up slice defaults
 }


 func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash:[]byte{}, Data:[]byte(data), PrevHash:prevHash}
	block.putHash()
	return block
}


func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{ []*Block{ Genesis() }}
}

func main() {

	chain := InitBlockChain()

	chain.AddBlock("first")
	chain.AddBlock("second")
	chain.AddBlock("third")

	for _, block := range chain.blocks {
		 fmt.Printf("-------------------\n")
		 fmt.Printf("Previous hash: %x\n", block.PrevHash)
		 fmt.Printf("data: %s\n", block.Data)
		 fmt.Printf("hash: %x\n", block.Hash)
	}

}