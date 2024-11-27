package task

import (
	"strconv"
	"strings"
)

// Реализация через срез
func EncryptThis(text string) string {

	if text == "" {
		return ""
	}

	wordL := strings.Split(text, " ")

	encryptL := make([]string, 0)

	for _, word := range wordL {
		if len(word) < 3 {
			firstRuneNumber := strconv.Itoa(int(rune(word[0])))
			encryptL = append(encryptL, firstRuneNumber+word[1:])
		} else {
			runes := []rune(word)

			firstRuneNumber := strconv.Itoa(int(runes[0]))

			runes[1], runes[len(runes)-1] = runes[len(runes)-1], runes[1]
			r := firstRuneNumber + string(runes[1:])
			encryptL = append(encryptL, r)
		}
	}

	return strings.Join(encryptL, " ")
}
