func trap(height []int) int {
    if len(height) == 0 {return 0} 
    res := 0
    l, r  := 0, len(height) - 1
    mL, mR := height[l], height[r]

    for l < r {
        if mL < mR {
            l++ 
            mL = max(mL, height[l])
            res += mL - height[l]
        } else { 
            r--
            mR = max(mR, height[r])
            res += mR - height[r]
        }
    }

    return res
}   

func max(a, b int) int {
    if a > b {return a} 
    return b
}