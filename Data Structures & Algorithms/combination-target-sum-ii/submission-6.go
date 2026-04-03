func combinationSum2(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var back func([]int, int, int)
	back = func(set []int, start, tar int) {
		if tar == target {
			add := make([]int, len(set))
			copy(add, set)
			res = append(res, add)
			return
		}  else if tar > target || start >= len(nums) {return}

		for i:=start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			back(set, i+1, tar+nums[i])
			set = set[:len(set)-1]
		}
	}

	back(make([]int,0),0,0)
	return res
}
