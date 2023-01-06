package main

import "fmt"

// 升级版

func main() {
	//fmt.Println(len("000000100000000000000000000000000000000000000000000000000000000000"))
	//block := NewBlock([]byte{}, "Bronson transfer Wendy 500 dollars")

	chain := newBlockChain()

	// 添加新区块
	chain.addBlockToChain("Bronson Transfer Wendy 5000 Big butt coin")
	chain.addBlockToChain("Bronson Transfer Wendy 10000 Big butt coin")

	for i, block := range chain.blocks {
		fmt.Printf("==== current block height %d\n====", i)
		fmt.Printf("previous blcok hash value : %x\n", block.PrevHash)
		fmt.Printf("current blcok hash value : %x\n", block.Hash)
		fmt.Printf("current block data ： %s\n", block.Data)

	}

}
