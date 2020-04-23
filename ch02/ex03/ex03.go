package main

import (
	"fmt"
	"github.com/oyama1/go_training/ch02/ex03/popcount"
)

func main() {
	fmt.Printf("count.go %d ", popcount.PopCount(11111))
	fmt.Printf("count.go %d ", popcount.PopCountForLoop(11111))
}