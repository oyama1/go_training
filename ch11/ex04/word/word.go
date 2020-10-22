package word

import "unicode"

// IsPalindrome 回文かどうかを判断する
func IsPalindrome(s string) bool {
	var letters []rune // rune(文字コード)配列

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r)) // unicodeの文字列なら、変換して詰める、変換して回文の判断を正確にしたい
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
