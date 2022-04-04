package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transaction  []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

}

func (b *Block) Print() {
	fmt.Printf("timestamp:        %d\n", b.timestamp)
	fmt.Printf("nonce:        %d\n", b.nonce)
	fmt.Printf("previous_hash:        %x\n", b.previousHash)
	fmt.Printf("transactions:        %s\n", b.transaction)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      ` json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transaction  []string ` json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transaction:  b.transaction,
	})
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	b := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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
	//blockChain.Print()
	previousHash := blockChain.LastBlock().Hash()

	blockChain.CreateBlock(5, previousHash)
	//blockChain.Print()
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)

	blockChain.Print()

}
