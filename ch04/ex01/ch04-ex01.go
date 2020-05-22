package main

import (
	"fmt"
	"crypto/sha256"
)

const (
	a = 1
	b // 指定がなければ、前と同じ値が代入される
	c
	d = 2
)

type Weekday int // 型の宣言
const (
	Sunday Weekday = iota // iota = 定数生成器 ゼロからincrementする、演算も可能
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 定数 = p84 実行時ではなく、コンパイル時に評価が行われることが保証されている式 = コンパイラが値を知っている状態
// _ = math.Pow(1000,iota) でエラーが発生する
// これはいけるiota + iota //これはいけない なぜ？ math.Pow(1000,iota + 1)
// ./main.go:28:2: const initializer math.Pow(1000, 0) is not a constant が発生する、どのタイミングでconstantになる？
// 実行時に処理が走るから？確認方法は？

// 数値 => constに変換できる？
// パラメータ => enum に変換して取り扱いたい

// TODO コンパクトには、1000 * 1000 .... でなければ良い？
const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000 // TODO overflow して無理
	YB = ZB * 1000
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB)

	c1:=sha256.Sum256([]byte("x")) // 64 byte
	fmt.Printf("%x\n", c1)
	fmt.Println(len("2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881"))
}
