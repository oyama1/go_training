package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// struct
type Person struct {
	name string
	age  int
}

func main() {
	// Display("test", 100)
	// Display("struct", Person{name: "John", age: 20})

	fmt.Println("キーが構造体か配列であるマップの表示")
	Display("map[Person]int", map[Person]int{Person{name: "John", age: 20}: 1, Person{name: "Bob", age: 30}: 2})
	Display("map[[]string]]int", map[[2]string]int{[2]string{"JVM", "Java"}: 6})
}

// 表示処理 (public)
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// 表示処理 (private)
// 引数 path 表示する場合のキーになるstring, v = 表示したい値
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}

	// 構造体はNumFieldでフィールドの数分表示する
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	// Mapの場合、path[formatされたフィールド] = 値という表示を行う
	// mapKeysはマップのキーごとにreflect.Valueのスライスを返す
	// mapIndex(key) はkeyに対するインデックスを返す
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// 表示のフォーマットを行う処理
// displayから、mapのキーやdefaultに入った値のフォーマットに使用する
func formatAtom(v reflect.Value) string {
	// データ種別ごとに値をどう表示するかのフォーマットを行う
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
