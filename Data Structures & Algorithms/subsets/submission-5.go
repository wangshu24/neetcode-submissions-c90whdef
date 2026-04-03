func subsets(nums []int) [][]int {
    res := [][]int{}
    
    var back func([]int, int)
    back = func(temp []int, start int) {
  
			tmp := make([]int, len(temp))
        	copy(tmp, temp)
        	res = append(res, tmp)
	

        for i := start; i < len(nums); i++ {
            temp = append(temp, nums[i])
            back(temp, i+1)
            temp = temp[:len(temp)-1]
        }
    }
    
    back([]int{}, 0)
    return res
}