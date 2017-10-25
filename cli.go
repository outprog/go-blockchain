package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//
type CLI struct {
	bc *Blockchain
}

// 运行CLI
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("address", "", "Transaction address")
	getBalanceData := getBalanceCmd.String("address", "", "Transaction address")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}
	if getBalanceCmd.Parsed() {
		if *getBalanceData == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

// 添加区块
func (cli *CLI) addBlock(address string) {
	cli.bc.AddBlock([]*Transaction{NewCoinbaseTX(address, "")})
	fmt.Print("Success!")
}

// 获得余额
func (cli *CLI) getBalance(address string) {
	balance := 0
	UTXOs := cli.bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}
	fmt.Printf("balance of %s: %d\n", address, balance)
}

// 输出所有区块
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Transactions Hash: %s\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 打印使用方法
func (cli *CLI) printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("  addblock -address TRANSACTION ADDRESS // add an address to the blockchain")
	fmt.Println("  getbalance -address TRANSACTION ADDRESS // add an address to the blockchain")
	fmt.Println("  printchain // print all the blocks of the blockchain")
}

// 无效命令
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
