func topKFrequent(nums []int, k int) []int {
	nMap := make(map[int]int)
	for _,i := range nums {
		nMap[i]++
	}
	// fmt.Println(nMap)	
	freq := make([][2]int, 0)
	for num, cnt := range nMap {
		freq = append(freq, [2]int{cnt, num})
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i][0] > freq[j][0]
	})
	// fmt.Println(freq)

	res := make([]int, k)
	for i:= 0; i < k;i++ {
		res[i] = freq[i][1]
	}
	return res
}
