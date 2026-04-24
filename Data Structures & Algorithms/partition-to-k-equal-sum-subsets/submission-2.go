import ("slices")
func canPartitionKSubsets(nums []int, k int) bool {
	var mid int 
	for _, num := range nums {
		mid+=num
	}
	if mid % k != 0 {return false}
	mid = mid/k

	slices.Sort(nums)
	used := make([]bool, len(nums))

	var back func(int, int, int) bool
	back = func(k, sum, index int) bool {
		if k == 0 {
			return true
		}
		if sum == mid {
			return back(k-1, 0,0)
		}

		for i:=index; i <len(nums); i++ {
			if !used[i] && sum + nums[i] <= mid {
				used[i] = true
				if back(k, sum+nums[i], i+1) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}

	
	return back(k,0,0)
}

