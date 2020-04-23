package main

import (
	"fmt"

	"github.com/oyama1/go_training/ch02/ex01/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))


	c := tempconv.Celsius(50)
	fmt.Println(tempconv.CToK(c))
}