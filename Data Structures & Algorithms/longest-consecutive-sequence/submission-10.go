func longestConsecutive(nums []int) int {
    res := 0
    found := make(map[int]struct{})
    for _,i := range nums {
        found[i] = struct{}{}
    }

    for _,i := range nums {
        if _,ok := found[i-1]; !ok {
            tmp := 1
            for { 
                if _,ok := found[i+tmp]; ok {
                    tmp++
                } else { break }
            }
            if tmp > res {
                res = tmp
            }
        }
    }
    return res
}
