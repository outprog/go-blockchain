package main

import (
	"fmt"
	//"strconv"
)

func main() {
	fmt.Println("init blockchain")
	bc := NewBlockChain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
