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
		fmt.Println("min is: ", min, " prof is: ", prof, " on iter: ", i)
	}

	return prof
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}