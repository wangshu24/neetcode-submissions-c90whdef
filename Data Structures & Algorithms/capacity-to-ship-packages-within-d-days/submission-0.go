import ("slices")

func shipWithinDays(weights []int, days int) int {
	minw := slices.Max(weights)
	maxw := 0
	for _, w := range weights {
		maxw+=w
	}

	fmt.Println(minw, maxw)
	for minw < maxw {
		mid := minw + (maxw - minw)/2
		day, cargo := 1, 0
		for i:=0; i < len(weights); i++ {
			cargo += weights[i]
			if cargo <= mid {continue} else {
				cargo = 0 + weights[i]
				day++
			}
		}
		if day > days {
			minw = mid + 1
		} else {maxw = mid}
	}


	return minw
}
