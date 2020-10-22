#!/bin/bash
cd `dirname $0`

# png => jpgへ変換
go run ex01.go -format=jpg < man.png > man_conv.jpg
# png => gifへ変換
go run ex01.go -format=gif < man.png > man_conv.gif
