package main

import (
	"fmt"
	"image/color"
	"math"
	"flag"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

// 色の定義
var (
	white = color.RGBA{0,0,0,255}
	whiteStr = "white"
	red = color.RGBA{255,0,0,255}
	redStr = "red"
	blue = color.RGBA{0,255,0,255}
	blueStr = "blue"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// フラグの定義
var (
	// go run ex03-01.go -funcFlag 0 という形式で実行する
	// 第二引数はデフォルト値
	functionFlag = flag.Int("funcFlag", 0, "function flag") 
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isAInf := corner(i+1, j)
			bx, by, isBInf := corner(i, j)
			cx, cy, isCInf := corner(i, j+1)
			dx, dy, isDInf := corner(i+1, j+1)

			// コーナーのどれかがInfならスキップ
			if !isAInf && !isBInf && !isCInf && !isDInf {
				// fill 属性で色を指定できる
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, whiteStr)
			}
		}
	}
	fmt.Println("</svg>")
}

// 浮動小数点数の結果を返す関数が失敗するかもしれない場合、失敗をbool戻り値で別に報告
// i,j マス目の角の座標である2つの値を返す
func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	// 2D格子のマスを意味する
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	// 3D関数の空間
	z := selectFunc(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	// 2D等角投影法の計算
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isInfinities(sx, sy)
}

func isInfinities (fx, fy float64) (bool) {
	// 無限大かどうかをチェックする math.IsInf関数
	// 第二引数 : > 0 なら正 / < 0 なら負 / = 0 なら両方、の無限大かチェック
	isXInf := math.IsInf(fx, 0)
	isYInf := math.IsInf(fy, 0)
	return isXInf || isYInf // true なら Inf
}

func selectFunc(x, y float64) float64 {
	flag.Parse() // フラグの読み込み
	switch *functionFlag {
		case 1 :
			return f1(x,y)
		case 2 :
			return f2(x,y)
		default :
			return f(x,y)
	}
}

func f(x, y float64) float64 {
	// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Abs(r / 20) // 式を変えてみる
}


func f2(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Round(r - (0.95 * r)) // ?
}

