package main
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	calcTime("echoCmdArgsByFor", echoCmdArgsByFor)
	calcTime("echoCmdArgsByRange", echoCmdArgsByRange)
	calcTime("echoCmdArgsByJoin", echoCmdArgsByJoin)
}

func echoCmdArgsByFor () {
	var s,sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // os.Argsはコマンドライン引数の配列
		sep = " "
	}
	fmt.Println(s)
}

func echoCmdArgsByRange () {
	s,sep := "",""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoCmdArgsByJoin() { 
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func calcTime(funcName string, fn func()) {
	start := time.Now()

	fn()// 引数の関数を実行

	secs := time.Since(start).Seconds()

	fmt.Println(fmt.Sprintf("funcName:%s takes %.10fs", funcName, secs))
}