func isAnagram(s string, t string) bool {
	sMap := make([]int, 26)
	tMap := make([]int, 26)

	for _, val := range s {
		sMap[val- 'a']++
	}
	for _, val := range t {
		tMap[val-'a']++
	}
	for i:=0; i < len(sMap); i++ {
		if sMap[i] != tMap[i] {return false}
	}

	return true
}
