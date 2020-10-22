package word

import (
	"math/rand"
	"testing"
	"time"
)

// exercise 11.3
func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25-2) + 2  // 1以下は回文になってしまうので、2 - 24の範囲を指定
	runes := make([]rune, n) // runeのスライスを定義

	// 最初と最後が別なら回文ではなくなる
	for i := 0; i < n; i++ {
		if i == 0 {
			runes[0] = rune(0x3041) // 適当な値
			continue
		} else if i == n-1 {
			runes[n-1] = rune(0x3042) // 最初と異なる適当な値
			continue
		}
		// 最初と最後以外はランダム
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
	}
	return string(runes)
}

// exercise 11.3 回文でないrandom文字列で、回文出ないことをテスト
func TestNotRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) { // 回文ならエラー
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

// random テスト from sample code
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// 前と後からの位置が同じ場所に同じ文字列をいれる
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
