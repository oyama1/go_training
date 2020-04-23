package popcount
// ベンチマークの実行 : go test -bench=.
// test サブコマンドで実行するので、ファイル名は、*_test.go としてその中にベンチマークメソッドを定義する

import (
	"testing"
)

// テストメソッド
// Testの接頭辞をつけ、引数に *testing.T のパラメータを設定する
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

// ベンチマークメソッド
// Benchmarkの接頭辞をつけ、引数に *testing.B のパラメータを設定する
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