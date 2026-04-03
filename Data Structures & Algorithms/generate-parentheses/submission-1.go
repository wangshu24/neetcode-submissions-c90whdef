func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var back func([]byte, int) 
	back = func(set []byte, count int) {
		if count == n*2 {
			if isWellform(string(set)) {
				str := string(set)
				res = append(res, str)
			} 
			return
		}

		set = append(set, '(')
		back(set, count+1)
		set = set[:len(set)-1]
		set = append(set, ')')
		back(set, count+1)
	}

	back([]byte{}, 0)
	return res
}

func isWellform(s string) bool {
	open := 0
	for _,val := range s {
		if val == '(' {
			open++
		} else {
			open--
		}
		if open < 0 {return false}
	}
	if open != 0 {return false}
	return true
}