package main

import (
	"fmt"
	"github.com/oyama1/go_training/ch02/ex04/popcount"
)

func main() {
	fmt.Printf("PopCount %d\n", popcount.PopCount(11111))
	fmt.Printf("PopCountForLoop %d\n", popcount.PopCountForLoop(11111))
	fmt.Printf("PopCount64times %d\n", popcount.PopCount64times(11111))
}