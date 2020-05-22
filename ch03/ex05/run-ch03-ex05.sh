#!/bin/sh
cd `dirname $0`
go run ch03-ex05.go > color-mandelbrot.svg
open .