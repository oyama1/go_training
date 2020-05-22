package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// 負ではない10真数表記整数文字列にカンマを挿入します
func comma(s string) string {
	var buffer bytes.Buffer // Bufferの定義

	// decimalを分離
	sIntAndDecimal := strings.Split(s, ".")
	sIntPart := sIntAndDecimal[0]

	n := len(sIntPart) // 文字列のサイズバイトスライスの数
	if n <= 3 {
		return s
	}

	var isFirstCommaFlag = false;
	// reverse がやりたいけど、標準パッケージにない？
	for index, w := range sIntPart { 
		// 文字列サイズ分 ループ / wにはRune値が入る
		// fmt.Fprintln(os.Stdout, isFirstCommaFlag,index,n)// Fprint = 出力先を明示できる、バッファメモリやos.Stdoutを第一引数にすれば標準出力に出力できる

		if shouldComma(isFirstCommaFlag, index, n) {
			// fmt.Fprintln(os.Stdout, w) 
			isFirstCommaFlag = true 
			buffer.WriteString(",") // 文字列をバッファへ書き込み
		}
		_, err := buffer.WriteRune(w)// Rune値をバッファへ書き込む

		if err != nil {
			// ErrTooLarge(バッファが大きくなりすぎで起きる？) の場合
			return buffer.String()
		}

	}

	// decimalがあれば付与
	if len(sIntAndDecimal) > 1 {
		buffer.WriteString("." + sIntAndDecimal[1])
	}

	return buffer.String()
	//return comma(s[:n-3]) + "," + s[n-3:] // 再帰
}

// commaを打つべきかをチェックする
func shouldComma(isFirstCommaFlag bool, index, length int) bool {
	if (index == 0) {
		return false
	}

	// commaを打つべきかの条件判断
	var shouldFirstComma = (!isFirstCommaFlag && (index == length % 3)) // まだうってない + 3で割ったあまりの値
	var shouldFirstCommaPlace = length % 3

	if shouldFirstComma {
		return true
	}

	// first以外
	var shouldComma = ((index - shouldFirstCommaPlace) % 3 == 0)
	if shouldComma {
		return true
	}
	return false
}

