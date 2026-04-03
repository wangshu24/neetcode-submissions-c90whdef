func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
	queue := make([]int, 0)
	l,r  := 0, 0
	for r < len(nums) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[r] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, r)

		if l > queue[0] {
			queue = queue[1:]
		}
	
		if r >= k - 1 {
			res = append(res, nums[queue[0]])
			l++
		}
		r++
	}

	return res
}
