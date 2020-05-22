#!/bin/sh
cd `dirname $0`

go run ch04-ex02.go x
go run ch04-ex02.go -sha384=true x
go run ch04-ex02.go -sha512=true x
go run ch04-ex02.go -sha384=true -sha512=true x
go run ch04-ex02.go -sha384=true -sha512=true x y