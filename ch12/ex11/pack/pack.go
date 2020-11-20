package pack

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

const portNumber = ":12345"

// 構造体 => パラメータに復元したURLを生成する
func Pack(ptr interface{}) string {
	v := reflect.ValueOf(ptr).Elem() // 構造体の変数を取得
	fmt.Printf("%v\n", v)

	buf := bytes.NewBufferString("")

	buf.WriteString("http://localhost" + portNumber + "/search?")
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField

		tag := fieldInfo.Tag    // a reflect.StructTag
		name := tag.Get("http") // タグの値を取得 = 実効的な値
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		paramKeyAndValue := fmt.Sprintf("%v=%v", name, v.Field(i))
		buf.WriteString(paramKeyAndValue)
		if i+1 < v.NumField() {
			buf.WriteByte('&')
		}
	}
	return buf.String()
}
