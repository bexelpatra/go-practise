package main

import (
	"fmt"
)

type point struct {
	x, y int
}

const (
	targetBits = 24
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("just simple msg")
	bc.AddBlock("nowdays i'm reading book written in english")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
	}
	fmt.Printf("%p\n", bc.blocks)

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Println(p)

}
