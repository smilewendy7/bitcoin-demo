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

// -- 给block

// 4. 引入区块链 ？？

//5.  添加区块

// 升级版

func main() {
	block := NewBlock([]byte{}, "Bronson transfer Wendy 500 dollars")
	fmt.Printf("前一区块是 %x\n", block.PrevHash)
	fmt.Printf("当前区块是 %x\n", block.Hash)
	fmt.Printf("区块数据 %s\n", block.Data)

}
