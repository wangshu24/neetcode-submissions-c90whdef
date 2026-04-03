func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {return false}
	var res bool
	s1arr := [26]byte{}
	s2arr := [26]byte{}
	for i:=0; i < len(s1); i++ {
		s1arr[s1[i] - 'a']++
		s2arr[s2[i] - 'a']++
	}
	fmt.Println(s1arr, s2arr)
	if s1arr == s2arr {return true}
	
	for i:= len(s1); i < len(s2); i++ {
		s2arr[s2[i-len(s1)] - 'a']--
		s2arr[s2[i] - 'a']++
		if s1arr == s2arr {return true}
	}
	

	return res
}
