package task

import (
	"sort"
	"strings"
)

type RomanConverter struct {
	nRim map[int]string
	keys []int
}

func NewRomanConverter(nRim map[int]string) *RomanConverter {
	keys := make([]int, 0, len(nRim))
	for k := range nRim {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	return &RomanConverter{nRim: nRim, keys: keys}
}

func (rc *RomanConverter) getNRim(n int) string {
	var result strings.Builder

	for _, k := range rc.keys {
		for n >= k {
			result.WriteString(rc.nRim[k])
			n -= k
		}
	}

	return result.String()
}

func (rc *RomanConverter) Solution(number int) string {
	var result strings.Builder
	// Данной функции splitNumber здесь быть не должно.
	numbersList := splitNumber(number)

	for _, n := range numbersList {
		result.WriteString(rc.getNRim(n))
	}

	return result.String()
}
