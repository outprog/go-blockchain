package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
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
	//b.SetHash()
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()

	b.Hash = hash[:]
	b.Nonce = nonce

	return b
}
