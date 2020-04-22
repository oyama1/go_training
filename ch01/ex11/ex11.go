package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	// main関数は一つのゴルーチン上で動作
	FetchAll()
}

func FetchAll() {
	if (len(os.Args) <= 1) {
		fmt.Println("set url as this command argument!!")
		return
	}
	start := time.Now()
	ch := make(chan string) // チャネル生成 (他ゴルーチンとの通信を行うオブジェクトのようなもの)
	for _, url := range os.Args[1:] {
		go fetch(AddHttpPrefix(url), ch) // 新たなゴルーチンを生成、そこで処理を行う
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) 
		// ch チャネルから受信、mainで走っているFetchAllが全て受け取る
		// ch チャネルと通信する複数のゴルーチンが ch チャネルを満たすとコールバックされる。満たされないとずっと待ちの状態になっている？
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func AddHttpPrefix(str string) string {	
	const http_schema, https_schema = "http://", "https://"

	prefixedStr := ""
	if (strings.HasPrefix(str, http_schema) || strings.HasPrefix(str, https_schema)) {
		prefixedStr = str
	} else {
		prefixedStr = http_schema + str
	}
	return prefixedStr
}

func fetch(url string, ch chan <- string) { // ch = channelにstringを渡すチャネル
	start := time.Now()

	// 応答が返ってこないケースの検証
	if (url == "http://sleep") {
		sleepProcess()
		// ch <- fmt.Sprint("sleep end") // これを追加するとちゃんと処理が停止 = チャネルが満たされないから、
		return // 何が起きるかというと、処理が終わらなくなる
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // チャネルへ値を出力
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	fetchProcessTime := fmt.Sprintf("Process time %.2fs : byte = %7d : url = %s", secs, nbytes, url)
	outputToFile(fetchProcessTime)
	ch <- fetchProcessTime
}

func outputToFile (message string) {
	fileNameToOutput := "processTimes.txt"
	createFile(fileNameToOutput)

    file, err := os.OpenFile(fileNameToOutput, os.O_WRONLY|os.O_APPEND, 0666) // 出力先を開く
    if err != nil {
		fmt.Fprintf(os.Stderr, "outputToFile error: %v\n", err)
    }
    defer file.Close() // defer = 遅延実行 、処理を関数の最後に実行する
    fmt.Fprintln(file, message) // ファイル書き込み
}

func createFile (fileName string) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
    	fmt.Println(fileName, " is created")
	}
}

func sleepProcess() {
	time.Sleep(time.Second * 30)
}