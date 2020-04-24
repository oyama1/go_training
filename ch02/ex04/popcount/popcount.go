package popcount

// pc[i]はiのポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int { // uint = 符号なし
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountForLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(uint(i)*8))])
	}
	return count
}

func CountBy64TimesShift(x uint64) int {
	count := 0                      // 最下位の桁数をクリアした数
	for i := uint(0); i < 64; i++ { // i に xを代入し、x分だけループします
		if x&(1<<i) != 0 {
			// 1<<i = << 左シフト、整数は2倍 / & = 論理積、ビットの桁同士が1のものだけ残る
			// 例 : 2(10) を 左へシフト = 4(100)へ、最下位が0になる = 1のフラグが立つものが左にずれていく
			// なので、元のxのビットに対して、1桁ずつ下の桁から論理積に欠けていく、これを64回繰り返す
			count++ // count を加算
		}
	}

	return count
}
