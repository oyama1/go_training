package main

import (
	"fmt"
	"os"
)

func main() {
	var str []string
	for _, arg := range os.Args[1:] {
		str = append(str, arg)
	}
	str = removeDup(str)
	fmt.Println(os.Args[1:], str)
}

// 隣接している重複を削除する
func removeDup(slice []string) []string {
	i := 0 // 重複をスキップした場合
	prev := ""

	for _, curr := range slice {
		
		//fmt.Println(i,prev,curr)
		
		if prev == curr {
			// skip
		} else {
			slice[i] = curr // 引数の基底配列を書き換える
			i++
		}
		prev = curr
	}

	return slice[:i] // 重複を削除した部分までのスライスを返す
}
