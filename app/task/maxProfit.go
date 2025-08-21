package task

// Task LeetCode MaxProfit

func MaxProfit1(prices []int) int {
	var profit int
	isActive := false
	var assets int

	for i, v := 0, 1; v < len(prices); i, v = i+1, v+1 {

		now, end := prices[i], prices[v]

		if !isActive {
			if now < end {
				assets = now
				isActive = true
			}
		} else {
			if now < end {
				continue
			} else {
				profit += now - assets
				isActive = false
			}
		}

		// Для последнего значения
		if v == len(prices)-1 && isActive {
			if now > end {
				profit += now - assets
				continue
			} else {
				profit += end - assets
			}
		}
	}

	return profit
}

func MaxProfit(prices []int) int {
	var profit int

	if len(prices) == 1 {
		return 0
	}

	for i, v := 0, 1; v < len(prices); i, v = i+1, v+1 {

		now, end := prices[i], prices[v]

		if v == len(prices) {
			if now < end {
				profit += end - now
			} else {
				continue
			}
		}

		if now < end {
			profit += end - now
		} else {
			continue
		}
	}

	return profit
}
