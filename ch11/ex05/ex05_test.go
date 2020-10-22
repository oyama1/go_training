package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":" // 各変数に各値を代入

	// テストデータを定義、struct で定義してそれにデータを詰める => テストケース追加はここに追加すれば良い
	var testInputs = []struct {
		str    string // 入力値
		sep    string // スプリットする記号
		expect int    // スプリットした後のサイズ
	}{
		{"a", ",", 1},
		{"a,b", ",", 2},
		{"a:b:c", ":", 3},
		{"a b c d", " ", 4},
		{"a/b/c/,d", "/", 4},
	}

	for _, input := range testInputs {
		words := strings.Split(input.str, input.sep) // sを sepの記号で分割する
		if got, want := len(words), input.expect; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
		}
	}
}
