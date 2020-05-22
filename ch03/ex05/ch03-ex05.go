package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2 //複素数平面
		width, height          = 1024, 1024     // スケール
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// 背景の色
			return black()
			//return color.Red{255 - contrast*n}
		}
	}
	// 集合図の色
	//return color.Black
	return acosRGBA(v)
}

func black() color.Color {
	black := color.RGBA{0, 0, 0, 0} // 型のコンストラクタに値を渡す場合は、type{val1,val2}
	return black
}

func red() color.Color {
	red := color.RGBA{255, 0, 0, 255}
	return red
}

// YCbCr の使い方 color.YCbCr{192, blue, red}

// interesting functionsを参考に処理を実装
func acosRGBA(z complex128) color.Color {
	v := cmplx.Acos(z) // 逆余弦の複素数　ふんわりとした認識
	red := uint8(imag(v)*128) + 127 // 虚数成分を赤
	blue := uint8(real(v)*128) + 127 // 実数成分を青
	green := uint8(0)
	return color.RGBA{red, green, blue, 255}
}

