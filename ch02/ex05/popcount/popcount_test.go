package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	input := uint64(1111)
	if (PopCount(input) == 6) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func TestPopCountForLoop(t *testing.T) {
	input := uint64(1111)
	if (PopCountForLoop(input) == 6) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func TestPopCount64times(t *testing.T) {
	input := uint64(1111)
	if (PopCount64times(input) == 6) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func TestPopCountClearBit_input_1111(t *testing.T) {
	input := uint64(1111)
	if (PopCountClearBit(input) == 6) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func TestPopCountClearBit_input_0(t *testing.T) {
	input := uint64(0)
	if (PopCountClearBit(input) == 0) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func TestPopCountClearBit_input_128(t *testing.T) {
	input := uint64(128)
	if (PopCountClearBit(input) == 1) {
		return // pass
	} else {
		t.Fatal("Test failed")
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountForLoop(uint64(i))
	}
}

func BenchmarkPopCount64times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64times(uint64(i))
	}
}

func BenchmarkPopCountClearBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearBit(uint64(i))
	}
}