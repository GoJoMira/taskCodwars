package task

func MoveZeroes1(nums []int) {
	idx := 0
	sumZero := 0
	for _, i := range nums {
		if i != 0 {
			nums[idx] = i
			idx++
		} else {
			sumZero++
		}
	}

	idx1 := len(nums) - 1
	for range sumZero {
		nums[idx1] = 0
		idx1--
	}
}

// LeetCode

func MoveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
