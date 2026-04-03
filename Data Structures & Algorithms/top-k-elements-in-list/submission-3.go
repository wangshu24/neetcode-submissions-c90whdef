func topKFrequent(nums []int, k int) []int {
	fMap := make(map[int]int)
	for _, val := range nums {
		fMap[val]++
	}

	fArr := make([][]int, len(nums)+1)
	for val, freq := range fMap {
		fArr[freq] = append(fArr[freq], val)
	}

	res := make([]int, 0)
	for i:= len(fArr)-1; i > 0; i-- {
		for _, val := range fArr[i] {
			res = append(res, val)
			if len(res) == k {
				return res
			}
		}
	}

	return res 
}
