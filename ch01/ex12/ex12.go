package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"strconv"
)

var mu sync.Mutex // 排他制御を行うロックオブジェクト
var count int

func main() {
	lissHandler := func (w http.ResponseWriter, r *http.Request) {
		// TODO クラス・構造体でセットした方が良い
	    cyclesParam, err := strconv.Atoi(r.FormValue("cycles"))
	    if err != nil {
	    	cyclesParam = 5
	    }
	    sizeParam, err := strconv.Atoi(r.FormValue("size")) 
	    if err != nil {
	    	sizeParam = 100
	    }
	    nframesParam, err := strconv.Atoi(r.FormValue("nframes")) 
	    if err != nil {
	    	nframesParam = 64
	    }
	    delayParam, err := strconv.Atoi(r.FormValue("delay")) 
	    if err != nil {
	    	delayParam = 8
	    }
		lissajous(w,cyclesParam,sizeParam,nframesParam,delayParam)
	}
	http.HandleFunc("/", lissHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{
		color.Black, 
		color.RGBA{0x00, 0xff, 0x00, 0xff}} 

const (
	blackIndex = 0
	greenIndex = 1 
)

func lissajous(out io.Writer, cyclesParam , sizeParam ,
 				nframesParam , delayParam int) {

	cycles := cyclesParam // 発振器 x が完了する集会の回数
	size := sizeParam // 画像キャンパスは [-size..+size] の範囲を扱う
	nframes := nframesParam // アニメーションフレーム数
	delay := delayParam // 10ms単位でのフレーム間の遅延
	
	const (res = 0.001) // 回転の分解能

	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes} // GIFオブジェクトに値を入れている
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0 ; t < float64(cycles)*2*math.Pi ; t += res {
			x:= math.Sin(t)
			y:= math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex) 
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意 : エンコードエラーを無視
}
