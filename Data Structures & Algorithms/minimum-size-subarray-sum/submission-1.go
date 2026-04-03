func minSubArrayLen(target int, nums []int) int {
	res, sum := math.MaxInt32, 0
	left := 0 

	for right := 0; right < len(nums); right++ {
		sum+=nums[right]
		for right >= left && sum >= target {
			sum-=nums[left]
			res = min(res, right - left + 1)
			left++
		}
	} 

	if res == math.MaxInt32 {res = 0}
	return res
}
