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

func PopCount64times(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(pc[byte(x>>(uint(i)*8))])
	}
	return count
}

func PopCountClearBit(x uint64) int {
	count := 0
	bitValue := x
	for bitValue > 0 {
		count ++ // ビット演算を行った回数
		bitValue = (bitValue&(bitValue-1)) // 1が立っている最下位のビットをクリア + クリア後の値をuint64に戻して返している。ちなみに := で入れるとループ内のローカル変数ができてしまう
	}
	return count
}