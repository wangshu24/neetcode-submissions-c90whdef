func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	prof, min := 0, prices[0]

	for i:= 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		prof = max(prof, prices[i]-min)
	}

	return prof
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}