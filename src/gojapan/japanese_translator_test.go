package gojapan

import (
	"fmt"
	"testing"
)

func TestTranslateKatakana(t *testing.T) {
	result := translate("happy")
	assertEquals(t, result.japanese, "ハッピー")
}

func TestTranslateKatakanaTwoWords(t *testing.T) {
	result := translate("panda box")
	assertEquals(t, result.japanese, "パンダボックス")
}

func TestTranslateKanji(t *testing.T) {
	result := translate("old")
	assertEquals(t, result.japanese, "古い")
}

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}