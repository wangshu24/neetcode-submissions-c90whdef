func letterCombinations(digits string) []string {
	if len(digits) == 0 {return []string{}}
	letters := make(map[byte][]byte)
	letters['2'] = []byte{'a','b','c'}
	letters['3'] = []byte{'d','e','f'}
	letters['4'] = []byte{'g','h','i'}
	letters['5'] = []byte{'j','k','l'}
	letters['6'] = []byte{'m','n','o'}
	letters['7'] = []byte{'p','q','r','s'}
	letters['8'] = []byte{'t','u','v'}
	letters['9'] = []byte{'w','x','y','z'}

	res := []string{}
	var back func([]byte, int)
	back = func(temp []byte, cur int) {
		if cur == len(digits) {
			res = append(res, string(temp))
			return 
		}

		chars := letters[digits[cur]]
		for _,val := range chars {
			temp = append(temp, val)
			back(temp, cur+1)
			temp = temp[:len(temp)-1]
		}
		return
	}

	back(make([]byte,0), 0)
	return res
}
