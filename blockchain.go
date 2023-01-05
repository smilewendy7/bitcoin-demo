package main

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
	genBlock := NewBlock([]byte{}, "====== The Big butt Coin is minted!!! ======")
	return genBlock

}

// 5.  添加区块到区块链上
func (blockChain *BlockChain) addBlockToChain(data string) {
	//a. 创建新的区块

	lastBlock := blockChain.blocks[len(blockChain.blocks)-1]

	prevBlockHash := lastBlock.Hash
	newBlock := NewBlock(prevBlockHash, data)
	//b. 添加到区块链中
	blockChain.blocks = append(blockChain.blocks, newBlock)

}
