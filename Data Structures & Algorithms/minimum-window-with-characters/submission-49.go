func minWindow(s string, t string) string {
	if t == "" || len(s) < len(t) {
		return ""
	}

	mapt := make(map[byte]int)

	for i:=0; i < len(t); i++  {
		mapt[t[i]]++
	}

	have, need := 0, len(mapt)
	res:= []int{-1, -1}

	resLen := math.MaxInt32
	l := 0
	window := make(map[byte]int)

	for r:=0; r < len(s); r++ {
		window[s[r]]++
		if mapt[s[r]] > 0 && window[s[r]] == mapt[s[r]] {
			have++
		}

		for have == need {
			if (r-l+1) < resLen {
				res = []int{l,r}
				resLen = r - l + 1
			}

			window[s[l]]--
			if mapt[s[l]] > 0 && window[s[l]] < mapt[s[l]] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}

	return s[res[0]:res[1]+1]
}
