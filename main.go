package main

import (
	"fmt"
	//"strconv"
)

func main() {
	fmt.Println("create blockchain")
	bc := NewBlockChain()
	fmt.Printf("LastHash: %x\n\n", bc.tip)
	bc.AddBlock("second blockchain")
	fmt.Printf("LastHash: %x\n\n", bc.tip)
	bc.AddBlock("third blockchain")
	fmt.Printf("LastHash: %x\n\n", bc.tip)
}
