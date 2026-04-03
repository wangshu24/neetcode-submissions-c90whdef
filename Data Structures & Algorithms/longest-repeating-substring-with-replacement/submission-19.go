func characterReplacement(s string, k int) int {
	var res int
	mask := make([]int, 26)
	tail := 0
	maxF := 0
	
	for i:=0; i < len(s); i++ {
		idx := s[i]-'A'
		mask[idx]++

		maxF = max(maxF, mask[idx])
		for i - tail - maxF + 1 > k {
			mask[s[tail]-'A']--
			tail++
		}
		
		if (i-tail+1) > res {
			res = i - tail + 1
		}
	}

	return res
}
