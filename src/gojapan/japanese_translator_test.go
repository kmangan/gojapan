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
	result := translate("banana split")
	assertEquals(t, result.japanese, "バナナスプリット")
}

func TestTranslateKanji(t *testing.T) {
	result := translate("old")
	assertEquals(t, result.japanese, "古い")
}

func TestTranslateKanjiTwoWords(t *testing.T) {
	result := translate("guardian angel")
	assertEquals(t, result.japanese, "守護天使")
}

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}