package task

func ReverseString(s []byte) {
	for i, j := 0, len(s); i == j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// LeetCode
// Почему меньше на одни 1000KB
func ReverseString1(s []byte) {

	var i, j int = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
