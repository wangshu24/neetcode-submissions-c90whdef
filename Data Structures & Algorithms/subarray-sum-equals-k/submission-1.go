func subarraySum(nums []int, k int) int {
	var res int
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	var tempSum int
	for _,val := range nums {
		tempSum+=val
		if count, there := prefixMap[tempSum - k]; there {
			res+=count
		}
		prefixMap[tempSum]++
	}

	return res
}
