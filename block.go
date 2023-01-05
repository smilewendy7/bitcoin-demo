package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// 第一个版本无工作量证明

// 1. 定义区块链结构
// 0） 定义结构 -
type Block struct {
	// 1. 版本号
	Version uint64

	// 2. 前区块哈希
	PrevHash []byte

	// 3. Merkel根 --- 一个哈希值
	MerkelRoot []byte

	// 4. 时间戳
	TimeStamp uint64

	// 5. 难度值

	Difficulty uint64

	//6. 随机数，就是挖矿要找的数据
	Nonce uint64

	// a. 当前区块哈希， 正常比特币区块链中没有当前区块的哈希， 我们为了方便做了简化
	Hash []byte

	// b. 数据
	Data []byte
}

/*
--- 进阶版本 ---
1. 补充区块字段
2. 更新计算机哈希值
3. 优化代码
*/

// 1） 前区块哈希
// 2）当前数据哈希
// 3） 数据

// 2. 创建区块？？？ - block 函数 -- 最后打印出来区块链
func NewBlock(prevBlockHash []byte, data string) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Hash:       []byte{}, // empty array
		Data:       []byte(data),
	}

	block.calHash()
	return &block

}

//3. 生成哈希 --

func (block *Block) calHash() {
	//var blockInfo []byte
	/*
		blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
		blockInfo = append(blockInfo, block.PrevHash...)
		blockInfo = append(blockInfo, block.MerkelRoot...)
		blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
		blockInfo = append(blockInfo, block.Data...)
	*/

	temp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(temp, []byte{})
	// 拼接所有数据
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// 实现一个辅助函数， 功能是将unit64转换成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
