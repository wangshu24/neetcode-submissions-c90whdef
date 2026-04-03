func twoSum(nums []int, target int) []int {
    prev := make(map[int]int)

	for idx, val := range nums {
		remain := target - val
		if ind, there := prev[remain]; there {
			return []int{ind, idx}
		}
		prev[val] = idx
	}

	return []int{}
}
