package main

import (
	"flag"
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	hash()
}

func hash() {
	// flag 渡し方 -sha256=true / -sha256 true
	var (
	    sha256flg = flag.Bool("sha256", true, "bool flag") // *sha256 のポインタに値が入る
	    sha384flg = flag.Bool("sha384", false, "bool flag")
	    sha512flg = flag.Bool("sha512", false, "bool flag")
	)

	flag.Parse() // flagの読み込み

	sha256val, sha384val, sha512val := *sha256flg, *sha384flg, *sha512flg

	// argsをひとつずつ変換処理をかける
	args := flag.Args()
	
	for _, arg := range args {
		fmt.Printf("arg =  %s\n", arg)
		
		if sha256val {
			fmt.Printf("sha256 %x\n", sha256.Sum256([]byte(arg)))
	    }

		if sha384val {
			fmt.Printf("sha384 %x\n", sha512.Sum384([]byte(arg)))
		}
		if sha512val {
			fmt.Printf("sha512 %x\n", sha512.Sum512([]byte(arg)))
		}
	}
}
