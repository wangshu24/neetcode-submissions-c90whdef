func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	left := 0
	maxFreq := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		count[idx]++
		
		// Track max frequency in window
		maxFreq = max(maxFreq, count[idx])

		// If more than k replacements needed, shrink
		for (right - left + 1) - maxFreq > k {
			count[s[left]-'A']--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
