package main
import (
	"fmt"
	"os"
)

func main() {
	echoCmdArgs()
}

func echoCmdArgs () {
	for index, arg := range os.Args[1:] {
		fmt.Println(fmt.Sprintf("%d %s", index, arg))
	}
}