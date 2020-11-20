package encode

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	var zeroInt int
	var zeroStr string
	var zeroSlice []string
	nonZeroSlice := []string{"nonZero"}

	var testInputs = []struct {
		value  interface{}
		expect interface{}
	}{
		{zeroInt, ""},
		{zeroStr, ""},
		{zeroSlice, ""},
		{nonZeroSlice, "(\"nonZero\")"},
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
