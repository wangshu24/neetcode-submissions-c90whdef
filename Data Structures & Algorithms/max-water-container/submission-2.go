func maxArea(nums []int) int {
    res := 0
    l,r  := 0, len(nums)-1
    
    for l < r {
        if nums[l] > nums[r] {
            if (nums[r] * (r-l)) > res {
                res = nums[r] * (r-l)
            }
            r--
        } else  {
            if (nums[l] * (r - l)) > res {
                res = nums[l] * (r - l)
            }
            l++
        }
    }

    return res
}
