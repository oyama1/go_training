package intset

import (
	"testing"
)

/*
 * tests added for ch11.ex02
 */

// 正常に追加されるかのテスト
func TestAddAndHas(t *testing.T) {
	// テストデータを定義、struct で定義してそれにデータを詰める
	var tests = []struct {
		input int // 入力値
	}{
		{1},
		{10},
		{100},
		{0},
	}

	var intSet IntSet
	intMap := map[int]bool{}
	// testデータを回して、対象メソッドの実行と結果を確認する
	for _, test := range tests {
		intSet.Add(test.input)
		intMap[test.input] = true
		if !intSet.Has(test.input) && !intMap[test.input] {
			t.Errorf("TestAddAndHas (%q)", test.input)
		}
	}
}

// 結合ができるかのテスト
func TestUnionWith(t *testing.T) {
	// テストデータを定義、struct で定義してそれにデータを詰める
	testSlice := [3]int{2, 3, 4}

	// 追加用データの用意
	var addSet IntSet
	addMap := map[int]bool{}
	for _, add := range testSlice {
		addSet.Add(add)
		addMap[add] = true
	}

	// 初期値に1をセット
	var intSet IntSet
	intSet.Add(1)
	intMap := map[int]bool{}
	intMap[1] = true

	// 結合
	intSet.UnionWith(&addSet)
	for key, value := range addMap {
		intMap[key] = value
	}

	//結合した値が intSetの方にすべて入っているかを確認する
	for key, _ := range intMap {
		if !intSet.Has(key) {
			// 入っていなかったらエラー
			t.Errorf("TestUnionWith fail %q ", key)
		}
	}
}
