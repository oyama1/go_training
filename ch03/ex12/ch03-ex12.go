package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("enter 2 phrase for checking anagram")
		os.Exit(1)
	}

	phrase1, phrase2 := os.Args[1], os.Args[2]
	if checkAnagram(phrase1, phrase2) {
		fmt.Printf("anagram! [%s] [%s] \n", phrase1, phrase2)
	} else {
		fmt.Printf("NOT anagram [%s] [%s] \n", phrase1, phrase2)
	}
}

// anagramをチェックする
func checkAnagram(phrase1, phrase2 string) bool {
	// phrase1とphrase2の文字数が一致するか
	if len(phrase1) != len(phrase2) {
		return false
	}

	// phrase2の文字列すべてがphrase1に含まれているか
	// bytePhrase1 := []rune(phrase1) // ルーンスライスへ変換 Rune = コードポイントの型 / byte = byte単位の型
	for _, w := range phrase2 {
		//fmt.Fprintln(os.Stdout, i, " ", w)
		if !strings.ContainsRune(phrase1, w) {
			return false
		}
	}
	return true
}
