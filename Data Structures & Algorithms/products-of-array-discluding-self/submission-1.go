func productExceptSelf(nums []int) []int {
    n := len(nums)
    res, pref, suf := make([]int, n), make([]int, n), make([]int, n)
    
    pref[0], suf[n-1] = 1, 1
    
    for i:=1;  i < n; i++ {
        pref[i] = nums[i-1] * pref[i-1] 
    }
    for i:=n-2; i > -1; i-- {
        suf[i] = nums[i+1] * suf[i+1]
    }
    for i:=0; i < n; i++ {
        res[i] = pref[i] * suf[i]
    }

    return res
}
