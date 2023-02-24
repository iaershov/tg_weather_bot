package presenter

import "unicode"

func ToUpper(str string, pos int) string {
	r := []rune(str)
	r[pos] = unicode.ToUpper(r[pos])
	return string(r)
}
