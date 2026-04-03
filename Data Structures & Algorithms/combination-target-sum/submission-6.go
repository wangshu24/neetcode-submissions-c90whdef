func combinationSum(nums []int, target int) [][]int {
    res := make([][]int, 0)

	var back func([]int, int, int)
	back = func(set []int, target, start int) {
		if target == 0 {
			add := make([]int, len(set))
			copy(add, set)
			res = append(res, add)
			return
		} else if target < 0 {return}
		
		for i:=start; i < len(nums); i++ {
			set = append(set, nums[i])
			back(set, target - nums[i], i)
			set = set[:len(set)-1]
		}
	}

	back(make([]int,0), target,0)
	return res
}