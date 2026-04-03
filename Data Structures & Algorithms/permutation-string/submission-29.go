func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {return false}
	
	tail, head:= 0, len(s1)-1
	var maps, mapl [26]int
	for i:=0; i < len(s1);i++ {
		maps[s1[i]-'a']++ 
		mapl[s2[i]-'a']++ 
	}
	if maps == mapl {
		return true
	}

	for i:=head+1; i < len(s2); i++ {
		mapl[s2[tail]-'a']--
		mapl[s2[i]-'a']++
		tail++
		if mapl == maps {
			return true
		}
	}
	
	return false
}
