package main
import (
	"fmt"
	"os" // importは複数記述もできる
	"strings"
)

func main() {
	echoCmdArgs()
}

func echoCmdArgs () {
	fmt.Println(strings.Join(os.Args[0:], " "))
}