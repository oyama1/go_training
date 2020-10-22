#!/bin/bash
cd `dirname $0`

# fetch(=ch1からコピー) meta name="go-import" に実際のホストサイトのリポジトリが記載してあるのでgrepで取得、本当にこれで良いのだろうか・・・？
go run fetch/fetch.go https://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import