func hasDuplicate(nums []int) bool {
    visited := make(map[int]struct{})

	for _, val := range nums {
		if _, there := visited[val]; !there {
			visited[val] = struct{}{}
		} else {return true}

	}
	return false
}
