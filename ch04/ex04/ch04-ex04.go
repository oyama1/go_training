package main

import (
	"fmt"
	"os"
)

func main() {
	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s) // "[0 1 2 3 4 5]"
	rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func rotate(s []int, position int) {
	fmt.Println("rotate")

	if position < 0 || len(s) < position {
		fmt.Println("this position is out of range")
		os.Exit(1)
	}

	// 所定の位置までの値をバッファへ詰める
	buf := s[:position]
	//fmt.Println(s, buf) // [0 1 2 3 4 5] [0 1]

	// 所定の位置からの後ろに、バッファに詰めた値を置く
	rotated := append(s[position:], buf...) // slice 同士をconcat : https://blog.golang.org/slices

	// reverse と同じく参照先を書き換える
	for i,_ := range s {
		s[i] = rotated[i]
	}
}

