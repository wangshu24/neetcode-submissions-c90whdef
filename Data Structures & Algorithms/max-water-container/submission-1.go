func maxArea(nums []int) int {
    res := 0
    l,r  := 0, len(nums)-1
    
    for l < r {
        if nums[l] > nums[r] {
            // fmt.Println(nums[r] * (r-l))
            if (nums[r] * (r-l)) > res {
                res = nums[r] * (r-l)
            }
            r--
        } else  {
            // fmt.Println(nums[l] * (r - l))
            if (nums[l] * (r - l)) > res {
                res = nums[l] * (r - l)
            }
            l++
        }
    }

    return res
}
