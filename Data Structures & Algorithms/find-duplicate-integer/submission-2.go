func findDuplicate(nums []int) int {
    var last int
	sort.Ints(nums)
	for _, val := range nums {
		if last != val {
			last = val
		} else {
			return val
		}
	} 

	return last
}
