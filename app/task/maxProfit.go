package task

// Task LeetCode MaxProfit

func MaxProfit(prices []int) int {
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
			if v == len(prices)-1 && isActive {
				if now > end {
					profit += now - assets
					continue
				} else {
					profit += end - assets
				}
			}

			if now < end {
				continue
			} else {
				profit += now - assets
				isActive = false
			}
		}
	}

	return profit
}
