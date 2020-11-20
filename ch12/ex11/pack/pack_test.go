package pack

import (
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {
	// テストデータを定義、struct で定義してそれにデータを詰める => テストケース追加はここに追加すれば良い
	var testInputs = []struct {
		Label  string `http:"l"`
		Result int    `http:"max"`
		Exact  bool   `http:"x"`
	}{
		{"a", 1, true},
		{"a", 1, false},
		{"", 1, false},
	}

	for _, input := range testInputs {
		packedUrl := Pack(&input) // sを sepの記号で分割する
		got := packedUrl
		want := fmt.Sprintf("http://localhost:12345/search?l=%v&max=%v&x=%v", input.Label, input.Result, input.Exact)
		if got != want {
			t.Errorf("Pack(%v) returned %v , want %v", input, got, want)
		}
	}
}
