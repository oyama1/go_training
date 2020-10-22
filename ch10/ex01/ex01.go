package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif" // パッケージにgifもあったので追加
	"image/jpeg"
	"image/png" // PNGデコーダの登録 = import時に init が呼び出される
	"io"
	"os"
)

func main() {
	formatFlag := getFormatFlag()
	decodeImage(os.Stdin, os.Stdout, formatFlag)
}

const (
	jpgFormat = "jpg"
	pngFormat = "png"
	gifFormat = "gif"
)

func getFormatFlag() string {
	var (
		format = flag.String("format", jpgFormat, "format") // -format=jpg の形式でパラメータを付与
	)
	flag.Parse()   // flag読み込み
	return *format // フラグ取得 -format=jpg で渡した値が、ポインタに帰ってくる
}

// decode処理、実行例 : go run ex01.go < man.png > man_conv.png
func decodeImage(in io.Reader, out io.Writer, formatFlag string) {
	if err := convert(in, out, formatFlag); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
}

// 画像変換変換処理 : error {} で処理結果を error で返す　成功ならnil
func convert(in io.Reader, out io.Writer, formatFlag string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	if formatFlag == jpgFormat {
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	} else if formatFlag == pngFormat {
		return png.Encode(out, img)
	} else if formatFlag == gifFormat {
		return gif.Encode(out, img, &gif.Options{NumColors: 256}) // NumColors = imageに使う画像色最大数 最大値の256をセット
	} else {
		return err
	}
}
