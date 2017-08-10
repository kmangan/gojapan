package gojapan

type translator interface {
	translate(u string) string
}

type translation struct {
	english, japanese string
}

func translate(s string) (translation) {
	result := translation{english:s, japanese:""}
	return result
}