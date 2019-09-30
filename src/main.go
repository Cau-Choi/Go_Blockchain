package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevBlock []byte
	hash      []byte
}

func (b *Block) SetHash() {
	h := sha256.New()
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var data []byte
	data = append(data, b.PrevBlock...)
	data = append(data, b.Data...)
	data = append(data, timestamp...) //input time now

	_, err := h.Write(data)
	if err != nil {
		log.Panic(err)
	}
	b.hash = h.Sum(nil)

}

//생성자 함수정의
func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlock, []byte("")}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte("0"))
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := &Block{time.Now().Unix(), []byte(data), prevBlock.hash, []byte("")}
	newBlock.SetHash()
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("send 1 BTC to ivan")

	bc.AddBlock("send 2 BTC to ivan")

	for _, block := range bc.blocks {
		fmt.Printf("%s\n", block.Data)
		fmt.Printf("%x\n", block.hash)

	}
}
