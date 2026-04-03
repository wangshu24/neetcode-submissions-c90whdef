func search(nums []int, target int) int {
    res := -1

    for ind,num := range nums {
        if num == target {
            return ind
        }
    }


    return res
}
