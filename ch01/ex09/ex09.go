package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	FetchUrls()
}

const (
	HTTP_SCHEMA = "http://"
)

func FetchUrls() {
	if (len(os.Args) <= 1) {
		fmt.Println("One more argument is needed")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		url := AddPrefix(arg, HTTP_SCHEMA)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("response status: %s", resp.Status)

		b, err := io.Copy(os.Stdout, resp.Body) // Stdout にそのまま流している
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading body %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println("response body: %s", b)
	}
}

func AddPrefix(str,prefix string) string {
	prefixedStr := ""
	if (!strings.HasPrefix(str, prefix)) {
		prefixedStr = prefix + str
	} else {
		prefixedStr = str
	}
	return prefixedStr
}