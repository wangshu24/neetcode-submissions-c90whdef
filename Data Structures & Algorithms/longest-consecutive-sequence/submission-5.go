func longestConsecutive(nums []int) int {
    res := 0
    store := make(map[int]struct{})
    for _, num := range nums {
        store[num] = struct{}{}
    }

    for _, num := range nums {
        streak, cur := 0, num
        for _,ok := store[cur]; ok; _,ok = store[cur] {
            streak++
            cur++
        }
        if streak > res {
            res = streak
        }
    }

    return res
}
