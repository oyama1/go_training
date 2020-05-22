#!/bin/sh
cd `dirname $0`

go run ch03-ex12.go noPhrase
go run ch03-ex12.go phrase1 phrase2
go run ch03-ex12.go phrase1 phrase1
go run ch03-ex12.go pfasafwe phrase1
go run ch03-ex12.go raseph1 phrase1
go run ch03-ex12.go phrase12 phrase1
go run ch03-ex12.go 日本語 phrase1
go run ch03-ex12.go 日本語 日本語
go run ch03-ex12.go 本語日 日本語
go run ch03-ex12.go こんにちわ わこんにち

