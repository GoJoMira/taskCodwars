package task

import (
	"strconv"
	"strings"
	"unicode"
)

func PlayPass(s string, n int) string {
	var result strings.Builder

	for i, r := range s {

		if unicode.IsLetter(r) {
			r = r + int32(n)
			if r > rune('z') {
				r -= rune('z')
			}
			if i%2 == 0 {
				result.WriteString(strings.ToUpper(string(r)))
				continue
			}
			result.WriteString(strings.ToLower(string(r)))
		} else if unicode.IsDigit(r) {
			digit := 9 - int(r-'0')
			result.WriteString(strconv.Itoa(digit))
		} else {
			result.WriteString(string(r))
		}
	}

	r := []rune(result.String())
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
