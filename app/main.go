package main

import (
	"helloapp/app/task"

	"math/rand"
	"strings"
	"time"
)

func usingBuilder(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}

func usingSlice(n int) string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, "a")
	}
	return strings.Join(result, "")
}

func usingCon(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += "a"
	}
	return result
}

func main() {

	// Устанавливаем сид для генерации случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Срез для хранения чисел
	numbers := make([]int, 10_000_000)

	// Генерация чисел
	for i := 0; i < 10_000_000; i++ {
		numbers[i] = rand.Intn(4000) // Генерирует числа от 0 до 3999
	}

	n := 1_000_000

	start := time.Now()
	usingBuilder(n)
	println("Builder:", time.Since(start).Milliseconds(), "ms")

	start = time.Now()
	usingSlice(n)
	println("Slice:", time.Since(start).Milliseconds(), "ms")

	// start = time.Now()
	// usingCon(n)
	// println("Con:", time.Since(start).Milliseconds()/1000, "sec")

	start = time.Now()
	for _, n := range numbers {
		task.Solution(n)
	}
	println("Solution:", time.Since(start).Milliseconds(), "ms")

	start = time.Now()
	for _, n := range numbers {
		task.Solution1(n)
	}
	println("Solution1:", time.Since(start).Milliseconds(), "ms")

	nRim := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// Создаем конвертер
	converter := task.NewRomanConverter(nRim)

	start = time.Now()
	for _, n := range numbers {
		converter.Solution(n)
	}
	println("Solution ChatGPT:", time.Since(start).Milliseconds(), "ms")
}
