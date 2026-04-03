func partition(s string) [][]string {
	res := [][]string{}

	var back func([]string,  int)
	back = func(temp []string, start int ){
		if start >= len(s) {
			add := make([]string, len(temp))
			copy(add, temp)
			res = append(res, add)
			return
		}
		
		for end:= start; end < len(s); end++ {
			if palindrome(s[start:end+1]) {
				temp = append(temp, s[start:end+1])
				back(temp, end+1)
				temp = temp[:len(temp)-1]
			}
		}
		
		return
	}

	back(make([]string, 0), 0)
	return res
}

func palindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {return true}
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {return false}
		l++
		r--
	}
	return true
}