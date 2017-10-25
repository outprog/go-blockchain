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

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	createBlockchainData := createBlockchainCmd.String("address", "", "Transaction adress")
	addBlockData := addBlockCmd.String("address", "", "Transaction address")

	switch os.Args[1] {
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if createBlockchainCmd.Parsed() {
		if *createBlockchainData == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainData)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

// 创建一个新的区块链
func (cli *CLI) createBlockchain(address string) {
	cli.bc.AddBlock([]*Transaction{NewCoinbaseTX(address, gensisCoinbaseData)})
	fmt.Print("Success!")
}

// 添加区块
func (cli *CLI) addBlock(address string) {
	cli.bc.AddBlock([]*Transaction{NewCoinbaseTX(address, "")})
	fmt.Print("Success!")
}

// 输出所有区块
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Transactions Hash: %x\n", block.HashTransactions())
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
	fmt.Println("  createblockchain -address TRANSACTION ADDRESS // add an address to the blockchain")
	fmt.Println("  addblock -address TRANSACTION ADDRESS // add an address to the blockchain")
	fmt.Println("  printchain // print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
