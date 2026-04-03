func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	var back func([]int,int)
	back = func(set []int, start int) {
		// if start == len(nums) {
			add := make([]int, len(set))
			copy(add, set)
			res = append(res, add)
		// 	return
		// }

		for i:=start; i  < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			back(set, i+1)
			set = set[:len(set)-1]
		}
	}
	back(make([]int,0), 0)
	return res
}
