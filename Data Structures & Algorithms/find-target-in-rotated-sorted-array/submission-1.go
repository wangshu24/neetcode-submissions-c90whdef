func search(nums []int, target int) int {
    res := -1
    for ind,i := range nums {
        if i == target {
            return ind
        }
    }
    return res
}
