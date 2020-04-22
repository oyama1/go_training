package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	readFile();
}

func readFile() {
	if (len(os.Args) <= 1) {
		fmt.Println("Set file name to read as this command arguments!!!")
		return
	}
	counts := make(map[string]int) // key = 重複行, value = 重複回数
	lineAndFileName := make(map[string]string)
	var files []string
	fileNames := os.Args[1:]

	for _, fileName := range fileNames {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup error: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			files = append(files, fileName) // filesに対して、fileNameを追加
			counts[line]++

			_, ok := lineAndFileName[line] // 重複行が常に検出済みかをチェック
			// ファイルにまたがる重複行を、lineをキーにfileNameをバリューにセット
			// TODO ネストをなんとかしたい
			if ok  {
				if fileName == lineAndFileName[line] {
					// 同じファイルで重複行が検出された場合、何もしない
					continue
				} else {
					// 別のファイルで重複行が検出された場合、ファイル名を追加
					lineAndFileName[line] = lineAndFileName[line] + "/" + fileName
					continue
				}
			} else {
				lineAndFileName[line] = fileName
				continue
			}
			
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t[%s] \t %s \n", count, line, lineAndFileName[line])
		}
	}
}
