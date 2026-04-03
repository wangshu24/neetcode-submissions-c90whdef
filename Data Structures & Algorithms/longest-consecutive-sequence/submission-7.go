func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sort.Ints(nums)

    res := 0 
    cur, streak := nums[0], 0
    i:=0
    for i < len(nums) {
        if cur != nums[i] {
            cur = nums[i]
            streak = 0
        }

        for i < len(nums) && cur == nums[i] {
            i++
        }
        streak++
        cur++
        if streak > res {
            res = streak
        }
    }
    return res
}
