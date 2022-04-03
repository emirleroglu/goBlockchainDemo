package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transaction  []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

}

func (b *Block) Print() {
	fmt.Printf("timestamp:        %d\n", b.timestamp)
	fmt.Printf("nonce:        %d\n", b.nonce)
	fmt.Printf("previous_hash:        %s\n", b.previousHash)
	fmt.Printf("transactions:        %s\n", b.transaction)
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))

		block.Print()
	}

	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {

	blockChain := NewBlockChain()

	blockChain.CreateBlock(5, "hash 1")
	blockChain.CreateBlock(2, "hash 2")

	blockChain.Print()

}
