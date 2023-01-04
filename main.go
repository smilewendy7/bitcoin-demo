package main

import (
	"crypto/sha256"
	"fmt"
)

// 第一个版本无工作量证明

// 1. 定义区块链结构
// 0） 定义结构 -
type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

// 1） 前区块哈希
// 2）当前数据哈希
// 3） 数据

// 2. 创建区块？？？ - block 函数 -- 最后打印出来区块链
func NewBlock(prevBlockHash []byte, data string) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, // empty array
		Data:     []byte(data),
	}

	block.calHash()
	return &block

}

//3. 生成哈希 --

func (block *Block) calHash() {
	hashData := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(hashData)
	block.Hash = hash[:]
}

// 4. 引入区块链  - 一组区块们

// 定义区块链结构
type BlockChain struct {
	blocks []*Block
}

// 定义一个区块链
func newBlockChain() *BlockChain {
	return &BlockChain{
		// 第一块是创世块
		blocks: []*Block{GenesisBlock()},
	}

}

// 第一块- 创世块
func GenesisBlock() *Block {
	//NewBlock(prevBlockHash []byte, data string)*Block
	genBlock := NewBlock([]byte{}, "韭菜币的诞生")
	return genBlock

}

//5.  添加区块
//a. 创建新的区块
//b. 添加到区块链中

// 升级版

func main() {
	//block := NewBlock([]byte{}, "Bronson transfer Wendy 500 dollars")

	chain := newBlockChain()

	for i, block := range chain.blocks {
		fmt.Printf("==== 当前层高为 %d\n====", i)
		fmt.Printf("前一区块哈希值是 %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值是 %x\n", block.Hash)
		fmt.Printf("区块数据 ： %s\n", block.Data)

	}

}
