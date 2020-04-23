package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	HTTP_SCHEMA = "http://"
)

func main() {
	FetchUrls()
}

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

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("response body: %s", b)
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