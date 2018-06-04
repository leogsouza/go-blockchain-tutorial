package main

import "fmt"

func main() {

	bc := NewBlockchain()

	bc.AddBlock("Send 0.01 BTC to Leo")
	bc.AddBlock("Send 0.11 BTC to Leo")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
