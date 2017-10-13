package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// 区块Hash
func (b *Block) SetHash() {
	headerArr := [][]byte{b.PrevBlockHash, b.Data, []byte(strconv.FormatInt(b.Timestamp, 10))}
	headers := bytes.Join(headerArr, []byte{})
	sha := sha256.Sum256(headers)
	b.Hash = sha[:]
}

// 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	b := &Block{}
	b.Timestamp = time.Now().Unix()
	b.Data = []byte(data)
	b.PrevBlockHash = prevBlockHash
	b.SetHash()
	return b
}

// 区块链结构
type BlockChain struct {
	blocks []*Block
}

// 创建区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewBlock("first block", []byte{})}}
}

// 添加一个区块链
func (bc *BlockChain) AddBlock(data string) {
	prevBc := bc.blocks[len(bc.blocks)-1]
	bc.blocks = append(bc.blocks, NewBlock(data, prevBc.Hash))
}

func main() {
	fmt.Println("create blockchain")
	bc := NewBlockChain()
	bc.AddBlock("second blockchain")
	bc.AddBlock("third blockchain")

	for _, block := range bc.blocks {
		fmt.Printf("timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
