func minEatingSpeed(piles []int, h int) int {
	res, min, max := 0, 1, math.MinInt
	
	for _,p := range piles {
		if p > max {
			max = p
		}
	}

	for min <= max {
		mid := min + (max-min)/2
		time := 0
		for _,p := range piles {
			
				time += int(math.Ceil(float64(p) / float64(mid)))
			
		}
		if time <= h  {
			res = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return res
}
