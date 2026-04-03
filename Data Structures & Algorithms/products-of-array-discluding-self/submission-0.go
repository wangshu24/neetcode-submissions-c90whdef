func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    for i,_ := range nums {
    res[i] = 1
        for j:=0; j < len(nums); j++ {
            if j != i {
                res[i] = res[i] * nums[j]
            } 
        }
    }
    return res
}
