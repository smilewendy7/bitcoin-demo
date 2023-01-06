package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义proof of work 结构

type ProofOfWork struct {
	block *Block
	// 目标值
	target *big.Int
}

// 2. 创建POW函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 我们指定的那渎职是一个string类型
	// *** 62个值以上才能挖矿 -- target 64 个值以上才快
	targetStr := "000000100000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow

}

func (pow *ProofOfWork) Run() (uint64, []byte) {
	// 得到nonce 和 hash
	//1. 拼装数据 （区块的数据， 还有不断变化的随机数）
	//将二维的切片数组连接起来， 返回一个一维切片
	//2. 做哈希运算
	// 3. 与pow中的target 进行比较
	// a. 找到了， 退出返回
	// b. 没找到，继续找，随机数+1

	block := pow.block
	var nonce uint64
	var hash [32]byte

	for {
		//1. 拼装数据 （区块的数据， 还有不断变化的随机数）
		//将二维的切片数组连接起来， 返回一个一维切片
		temp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(temp, []byte{})

		//2. 做哈希运算
		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}

		// tepInt 是hash 的 big int 值
		tmpInt.SetBytes(hash[:])

		// 3. 与pow中的target 进行比较
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("===== Mining sucess!!！===== \n The nonce number is : %d\n, The hash number is : %x\n", nonce, hash)
			return nonce, hash[:]
		} else {
			nonce++ // nonce += 1
		}

	}

}
