func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	
	for i:=0; i < len(nums);i++ {
		remain := target  - nums[i] 
		diffs[remain] = i
	}
	fmt.Println(diffs)
	for i:=0; i < len(nums); i++ {
		idx, there := diffs[nums[i]]
		
		if there && i != idx {
			fmt.Println(nums[i], nums[idx], idx, i)
			return []int{i, idx}
		}
	}

	return []int{}
}
