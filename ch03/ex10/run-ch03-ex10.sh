#!/bin/sh
cd `dirname $0`

go run ch03-ex10.go 100
go run ch03-ex10.go 100.12
go run ch03-ex10.go 100000.12
go run ch03-ex10.go 100 10000 1000000
go run ch03-ex10.go 100.12 10000.12 1000000.12

