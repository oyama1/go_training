package encode

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	// テストデータを定義、struct で定義してそれにデータを詰める => テストケース追加はここに追加すれば良い
	var testInputs = []struct {
		value  interface{}
		expect interface{}
	}{
		{"a", "\"a\""},
		{false, "nil"},
		{10i, "#C(0.0 10.0)"},
		{float64(10.0), "10.000000"},
		{[]int{1, 2, 3}, "(1 2 3)"},
	}

	for _, input := range testInputs {
		buf, _ := Marshal(input.value)
		got, want := string(buf), input.expect
		if got != want {
			t.Errorf("Marshal(%v) returned %v , want %v", input.value, got, want)
		}
		fmt.Printf("Marshal %v to %v", input.value, want)
	}
}
