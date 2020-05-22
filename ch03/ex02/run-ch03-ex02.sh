#!/bin/sh
cd `dirname $0`
go run ch03-ex02.go > polygon.svg 
go run ch03-ex02.go -funcFlag=1 > polygon-f1.svg 
go run ch03-ex02.go -funcFlag=2 > polygon-f2.svg 
