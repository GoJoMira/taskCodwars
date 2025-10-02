package task

func RotateImage(matrix [][]int) {
	for i := 0; i < len(matrix)-1; i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		for _, row := range matrix {
			row[i], row[j] = row[j], row[i]
		}
	}
}

// LeetCode
func RotateImage1(m [][]int) {
	n := len(m)
	// mirror along diagonal
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	// mirror left-right
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			m[i][j], m[i][n-j-1] = m[i][n-j-1], m[i][j]
		}
	}
}
