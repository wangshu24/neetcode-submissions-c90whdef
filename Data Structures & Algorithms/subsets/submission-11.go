func subsets(nums []int) [][]int {
    res := [][]int{}
    
    var back func([]int, int)
    back = func(temp []int, start int) {
		if start >= len(nums) {
			add := make([]int, len(temp))
			copy(add, temp)
			res = append(res, add)
			return
		}
		temp = append(temp, nums[start])
		back(temp, start+1)
		temp = temp[:len(temp)-1]
		back(temp, start+1)
    }
    
    back([]int{}, 0)
    return res
}