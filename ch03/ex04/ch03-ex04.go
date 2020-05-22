package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

var (
	width, height = 600.0, 320.0            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	colorValue    = "white"
)
// フラグの定義
var (
	// go run ex03-01.go -funcFlag 0 という形式で実行する
	// 第二引数はデフォルト値
	functionFlag = flag.Int("f", 0, "function flag") 
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var mu sync.Mutex // 排他制御を行うロックオブジェクト
var count int

func main() {
	fmt.Println("===== Server Running =====")
	fmt.Println("server running port 8000 can handle path [/svg] and parameters height, width, color")

	svgHandler := func(w http.ResponseWriter, r *http.Request) {
		// svg であることをヘッダーで示す
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w, r)
	}
	http.HandleFunc("/svg", svgHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svg(out io.Writer, r *http.Request) {
	// 高さと幅をパラメータから取得
	heightParam, err := strconv.ParseFloat(r.FormValue("height"), 64)
	if err != nil {
		fmt.Println("heightParam not set ")
		height = 320.0
	} else {
		height = heightParam
	}

	widthParam, err := strconv.ParseFloat(r.FormValue("width"), 64)
	if err != nil {
		fmt.Println("widthParam not set ")
		width = 600.0
	} else {
		width = widthParam
	}

	// スケールを再計算
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit

	//  色をパラメータから取得
	colorParam := r.FormValue("color")
	if colorParam != "" {
		fmt.Println("colorParam not set ")
	} 
	colorValue = colorParam

	cornerSize := 1

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	//fmt.Fprintf(out, "test")
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isAInf := corner(i+cornerSize, j)
			bx, by, isBInf := corner(i, j)
			cx, cy, isCInf := corner(i, j+cornerSize)
			dx, dy, isDInf := corner(i+cornerSize, j+cornerSize)
			if !isAInf && !isBInf && !isCInf && !isDInf {
				// fill 属性で色を指定できる
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",ax, ay, bx, by, cx, cy, dx, dy, colorValue)
// fmt.Fprintf(out, "test")
			}
		}
	}
	fmt.Fprintf(out,"</svg>")
}

// 浮動小数点数の結果を返す関数が失敗するかもしれない場合、失敗をbool戻り値で別に報告
// i,j マス目の角の座標である2つの値を返す
func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	// 2D格子のマスを意味する
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	// 3D関数の空間
	z := selectFunc(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	// 2D等角投影法の計算
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isInfinities(sx, sy)
}

func isInfinities(fx, fy float64) bool {
	// 無限大かどうかをチェックする math.IsInf関数
	// 第二引数 : > 0 なら正 / < 0 なら負 / = 0 なら両方、の無限大かチェック
	isXInf := math.IsInf(fx, 0)
	isYInf := math.IsInf(fy, 0)
	return isXInf || isYInf // true なら Inf
}

func selectFunc(x, y float64) float64 {
	flag.Parse()
	switch *functionFlag {
	case 1:
		return f1(x, y)
	case 2:
		return f2(x, y)
	default:
		return f(x, y)
	}
}

func f(x, y float64) float64 {
	// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Abs(r / 20)
}

func f2(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Round(r - (0.95 * r)) // ?
}
