package main

import (
	"fmt"
)

func main() {
	//!+array
	b := [...]int{0, 1, 2, 3, 4, 5}
	reversePointer(&b) // ポインタ *b を生成して渡す
	fmt.Println(b)     // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev sample code
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] // 最初の位置と最後の位置を順に入れ替えていく
	}
}

//!-rev

// *[]type でtypeのポインタをさす
// 配列は [num] で配列のサイズが決まらないとできない
// 可変長は配列？スライスで賄えるので、わざわざ配列を可変にする方法は考えなくて良い
// reverse(sp *[6]int, num int) => 同じメソッド名は無理かも
// ./main.go:34:9: not enough arguments in call to reverse
//	have ([]int)
//	want (*[6]int, int)
func reversePointer(sp *[6]int) {
	for i, j := 0, len(sp)-1; i < j; i, j = i+1, j-1 { // i = start, j = end
		// s[i] と s[j] の文字を入れ替える
		sp[i], sp[j] = sp[j], sp[i]
	}
}
