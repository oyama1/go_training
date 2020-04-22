package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
		color.Black, 
		color.White, 
		color.RGBA{0xff, 0x00, 0x00, 0xff}, 
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff}} 

const (
	blackIndex = 0
	whiteIndex = 1 
	redIndex = 2
	greenIndex = 3
	blueIndex = 4
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // 発振器 x が完了する集会の回数
		res = 0.001 // 回転の分解能 
		size = 100 // 画像キャンパスは [-size..+size] の範囲を扱う
		nframes = 64 // アニメーションフレーム数
		delay = 8 // 10ms単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes} // GIFオブジェクトに値を入れている
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ { // 64フレームで描画しているのか、64回回って画像を描画している

		colorIndex := RandomIndex()

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0 ; t < cycles*2*math.Pi ; t += res {
			x:= math.Sin(t)
			y:= math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex) 
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意 : エンコードエラーを無視
}

// 色のランダム生成メソッド
func RandomIndex() uint8 {
	variation := 4 // 生成する色の種類
	offset := 1 // paletteの何個目の色から使用するか
    return uint8(rand.Intn(variation) + offset)
}

