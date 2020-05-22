#!/bin/sh
cd `dirname $0`

go run ch04-ex05.go a b c d
go run ch04-ex05.go ab ab ab abc abc abcd
go run ch04-ex05.go 1 1 2 3 3 2