package main

import (
	"fmt"
	"log"
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

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func (b *Block) print() {
	fmt.Printf("timestamp			%d", b.timestamp)
	fmt.Printf("nonce				%d", b.nonce)
	fmt.Printf("previousHash		%s", b.previousHash)
	fmt.Printf("transactions		%s", b.transaction)
}
func main() {
	blockchain := NewBlock()
	b.print()
}
