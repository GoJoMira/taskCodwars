package task

import (
	"strings"
)

func IsValidSudoku1(board [][]string) bool {

	m := make(map[string]struct{})

	for _, row := range board {
		for _, v := range row {
			if v == "." {
				continue
			}
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
			} else {
				return false
			}
		}
		m = make(map[string]struct{})
	}

	m = make(map[string]struct{})

	for i := range len(board) - 1 {
		for _, column := range board {
			v := column[i]
			if v == "." {
				continue
			}
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
			} else {
				return false
			}
		}
		m = make(map[string]struct{})
	}

	m = make(map[string]struct{})

	for i := 0; i < len(board); i += 3 {
		x := board[i : i+3]
		for y := 0; y < len(board); y += 3 {
			for _, v := range x {
				for _, v1 := range []string{v[y], v[y+1], v[y+2]} {
					if v1 == "." {
						continue
					}
					if _, ok := m[v1]; !ok {
						m[v1] = struct{}{}
					} else {
						return false
					}
				}
			}
			m = make(map[string]struct{})
		}
	}

	return true
}

func IsValidSudoku2(board [][]byte) bool {

	m := make(map[byte]struct{})

	for _, row := range board {
		for i, v := range row {
			_ = i
			if v == 46 {
				continue
			}
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
			} else {
				return false
			}
		}
		m = make(map[byte]struct{})
	}

	m = make(map[byte]struct{})

	for i := range len(board) {
		for _, column := range board {
			v := column[i]
			if v == 46 {
				continue
			}
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
			} else {
				return false
			}
		}
		m = make(map[byte]struct{})
	}

	m = make(map[byte]struct{})

	for i := 0; i < len(board); i += 3 {
		x := board[i : i+3]
		for y := 0; y < len(board); y += 3 {
			for _, v := range x {
				for _, v1 := range []byte{v[y], v[y+1], v[y+2]} {
					if v1 == 46 {
						continue
					}
					if _, ok := m[v1]; !ok {
						m[v1] = struct{}{}
					} else {
						return false
					}
				}
			}
			m = make(map[byte]struct{})
		}
	}

	return true
}

func StringSliceToByte(SS [][]string) [][]byte {
	SB := make([][]byte, len(SS))

	var strS = strings.Builder{}
	for idx, v := range SS {
		for _, i := range v {
			strS.WriteString(i)
		}
		SB[idx] = []byte(strS.String())
		strS.Reset()
	}

	return SB
}

// Version LeetCode

func IsValidSudoku(board [][]byte) bool {
	var row [9][9]bool
	var col [9][9]bool
	var cell [9][9]bool

	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				continue
			}

			val := int(board[i][j] - '1')

			cellIndex := (i/3)*3 + (j / 3)

			if row[i][val] || col[j][val] || cell[cellIndex][val] {
				return false
			}

			row[i][val] = true
			col[j][val] = true
			cell[cellIndex][val] = true
		}
	}

	return true
}
