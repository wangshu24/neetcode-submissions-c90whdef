func permute(nums []int) [][]int {
	res := [][]int{}
	visit := make(map[int]struct{})

	var back func(map[int]struct{}, []int)
	back = func(vis map[int]struct{}, set []int) {
		if len(set) == len(nums) {
			add := make([]int, len(set))
			copy(add, set)
			res = append(res, add)
			return 
		}

		for i:=0; i < len(nums); i++ {
			if _,there := vis[nums[i]]; !there {
				set = append(set, nums[i])
				vis[nums[i]] = struct{}{}
				back(vis, set)
				delete(vis, nums[i])
				set = set[:len(set)-1]
			} else {continue}
		}
	}

	back(visit, make([]int,0))
	return res
}
