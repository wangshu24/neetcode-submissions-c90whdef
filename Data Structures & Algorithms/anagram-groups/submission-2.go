func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	visited := make(map[[26]byte][]string)

	for _, str := range strs {
		chars := [26]byte{} 
		for i:=0; i < len(str); i++ {
			chars[str[i]-'a']++
		}
		if _, there := visited[chars]; !there {
			visited[chars] = []string{}
		} 
		visited[chars] = append(visited[chars], str)
		
	}

	for _, anagram := range visited {
		res = append(res, anagram)
	}

	return res
}
