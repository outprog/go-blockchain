package main

import (
	"fmt"
	"strconv"
)

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
		fmt.Printf("Nonce: %d\n", block.Nonce)

		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
