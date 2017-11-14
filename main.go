package main

import (
	"fmt"
	//"strconv"
)

func main() {

	bc := NewBlockchain("15XjNnP5F5JxaqKuExyzmKaW8UqBePxodN")
	fmt.Println("load blockchain")
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()

}
