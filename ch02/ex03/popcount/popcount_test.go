package popcount
// ベンチマークの実行 : go test -bench=.
// test サブコマンドで実行するので、ファイル名は、*_test.go としてその中にベンチマークメソッドを定義する

import (
	"testing"
)

// ベンチマークメソッド
// Benchmarkの接頭辞をつけ、引数に*testing.B のパラメータを設定する
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