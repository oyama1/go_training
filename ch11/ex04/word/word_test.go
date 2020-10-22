package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// random テスト from sample code
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)

	// TODO 記号を自動生成したい
	var symbols = []rune{',', '、', '.', '。', ' '}
	randSymbolNum := rng.Intn(5)

	for i := 0; i < (n+1)/2; i++ {
		// 前と後からの位置が同じ場所に同じ文字列をいれる
		// r := rune(rng.Intn(0xFFFF)) // FFFFまででUnicode的には記号が範囲に入りそう？ => 出力でみなかったので明示して差し込む
		r := rune(rng.Intn(0x1000))
		if i == randSymbolNum {
			r = symbols[randSymbolNum] // TODO 特定の値の時記号を差し込んでいるが、もっとスマートに解決したい、対応する記号も少ない
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	fmt.Println(string(runes))
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
