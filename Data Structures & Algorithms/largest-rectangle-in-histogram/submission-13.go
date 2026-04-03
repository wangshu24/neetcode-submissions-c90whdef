type Area struct {
    pref int
    suf int
}

func largestRectangleArea(heights []int) int {
    var res int
    stack := make([]Area, len(heights))
    for i:=0; i < len(heights); i++ {
        l, r := 0,  0
        for i-l-1 >= 0 && heights[i] <= heights[i-l-1] {
            l++
        } 
        for i+r+1 < len(heights) && heights[i] <= heights[i+r+1] {
            r++
        }
        fmt.Println(l,r)
        stack[i] = Area{pref: l, suf: r}
    }

    for i, a := range stack {
        res = max(res, heights[i] * (a.pref + a.suf + 1))
    }

    return res
}

func max(a,b int) int{
    if a > b {return a}
    return b
}