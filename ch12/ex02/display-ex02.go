package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 循環参照を持つ構造体
type Cycle struct {
	Value int
	Tail  *Cycle
}

func main() {
	fmt.Println("再起呼び出しの確認 5階層まで表示")
	var c Cycle
	c = Cycle{42, &c}
	Display("cycle", c)

	// map
	// cMap := map[int]Cycle{1: c, 2: c, 3: c}
	// Display("cycleMap", cMap)
}

// ネストの表示制限
const maxNestedCount = 3

// 表示処理 (public)
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

func display(path string, v reflect.Value, currNestedCount int) {
	// fmt.Printf("nest count %d \n", currNestedCount)
	if currNestedCount > maxNestedCount {
		// ネストが一定数を超えた場合表示を終了する
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), currNestedCount)
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), currNestedCount)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), currNestedCount)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// 参照した数 = ネストの値とするため、現在の表示階層数に+1
			display(fmt.Sprintf("(*%s)", path), v.Elem(), currNestedCount+1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), currNestedCount)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// 表示のフォーマットを行う処理
// displayから、mapのキーやdefaultに入った値のフォーマットに使用する
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)

	// ...floating-point and complex cases omitted for brevity...
	// 小数点や複素数は後の練習問題で行う

	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)

	// 構造体(struct)の表示フォーマット:type{field:value}
	case reflect.Struct, reflect.Array:
		// 文字列を結合するためのbyte配列
		var stringByte []byte
		stringByte = append(stringByte, v.Type().String()...)    // append で 第一引数のbufferにstringByteに第二引数の値をつめて返す
		stringByte = append(stringByte, fmt.Sprintf("%v", v)...) // vで構造体の中身を取得、文字列... でbyteに変換できる
		// byte => 文字列へ戻す
		return string(stringByte)

	// 配列の表示フォーマット:

	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
