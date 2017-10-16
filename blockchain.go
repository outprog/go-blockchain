package main

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
