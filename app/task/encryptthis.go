package task

import (
	"strconv"
	"strings"
)

// Реализация через срез
func EncryptThis_1(text string) string {

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

// Убрал дублиющий код
func EncryptThis_2(text string) string {

	if text == "" {
		return ""
	}

	wordL := strings.Split(text, " ")
	encryptL := make([]string, 0)

	for _, word := range wordL {
		// Делаю каждый раз преобразование
		// даже когда мне это не нужно.
		w := []byte(word)
		// Можно было сразу сделать преобразование
		// в строку, а не делать это позже
		// firstR := strconv.Itoa(int(w[0]))
		firstR := w[0]

		if len(w) >= 3 {
			w[1], w[len(w)-1] = w[len(w)-1], w[1]
		}

		w = w[1:]
		encryptL = append(encryptL, strconv.Itoa(int(firstR))+string(w))
	}

	return strings.Join(encryptL, " ")

}

// Вариант ChatGPT
func EncryptThis_3(text string) string {
	if text == "" {
		return ""
	}

	wordL := strings.Split(text, " ")
	encryptL := make([]string, len(wordL))

	for i, word := range wordL {
		if len(word) == 0 {
			continue
		}

		firstChar := strconv.Itoa(int(word[0])) // Номер первого символа
		if len(word) == 1 {
			// Если слово состоит из одного символа
			encryptL[i] = firstChar
			continue
		}

		if len(word) == 2 {
			// Если слово состоит из двух символов
			encryptL[i] = firstChar + word[1:]
			continue
		}

		// Слово из 3 и более символов
		chars := []byte(word)
		chars[1], chars[len(chars)-1] = chars[len(chars)-1], chars[1] // Меняем второй и последний символы

		encryptL[i] = firstChar + string(chars[1:]) // Конкатенация
	}

	return strings.Join(encryptL, " ")
}

// Решение с CodeWars
// ChatGPT заявляет что case 1 and case len(word)-1
// избыточны для коротких слов, 1 или 2 символа
func EncryptThis(text string) string {
	res := ""
	for _, word := range strings.Split(text, " ") {
		for i, letter := range word {
			switch i {
			case 0:
				res += string(strconv.Itoa(int(letter)))
			case 1:
				res += string(word[len(word)-1])
			case len(word) - 1:
				res += string(word[1])
			default:
				res += string(letter)
			}
		}
		res += " "
	}
	return strings.TrimSpace(res)
}
