func longestConsecutive(nums []int) int {
    res := 0
    n := make(map[int]struct{})
    for _, i := range nums {
        n[i] = struct{}{}
    }

    for i,_ := range n {
        if _,ok := n[i-1]; !ok {    
            tmp, cur := i, 1
            for {
                if _,ok := n[tmp+1]; ok {
                    cur++
                    tmp++
                } else {
                    break
                }
            }
            if cur > res {
                res = cur
            }
        }
    }
    return res
}
