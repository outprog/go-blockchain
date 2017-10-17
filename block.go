package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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

// 序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
