package gojapan

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	result := translate("happy")
	assertEquals(t, result, "ハッピー")
}

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}