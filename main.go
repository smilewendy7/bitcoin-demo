package main

import (
	"crypto/sha256"
	"fmt"
)

// 0. 定义结构
type Block struct {
	//1. 前区块哈希
	PrevHash []byte
	//2. 当前区块哈希
	Hash []byte
	//3. 数据
	Data []byte
}

//二. 创建区块

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, // 先定义空， 再后边计算
		Data:     []byte(data),
	}

	block.SetHash()

	return &block

}

// 三. 生成哈希
// 给block 注册一个函数 ***
func (block *Block) SetHash() {
	// block 内部做hash
	// 1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	// 2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// 四. 引入区块链
// 把所有数据穿起来-数组
type BlockChain struct {
	//当以一个区块链数组
	blocks []*Block
}

// 五. 定义一个区块链
func NewBlockChain() *BlockChain {
	// 一般定义区块链要往里边塞上一块儿糖
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 创世块
func GenesisBlock() *Block {
	return NewBlock("GO一期创世块， 老牛逼了！", []byte{})

}

//五. 添加区块
//六. 重构代码

func main() {
	bc := NewBlockChain()
	for i, block := range bc.blocks {
		fmt.Printf("======= 当前前区块高度： %d======\n", i)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)

	}

}
