func lengthOfLongestSubstring(s string) int {
	visit := make(map[byte]struct{})
	tail, m := 0, 0

	for i:=0; i < len(s); i++ {
		for {
			if _,there := visit[s[i]]; there {
				delete(visit, s[tail])
				tail++
			} else {

			break
			}
		}
		m = max(m, i - tail + 1)
		visit[s[i]] = struct{}{}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}