func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers)-1
    res := make([]int, 2)
    for l < r {
        for r > l && numbers[r] + numbers[l] > target  {
            r--
        }
        for l < r && numbers[r] + numbers[l] < target {
            l++
        }
        if numbers[r] + numbers[l] == target {
            res[0] = l+1
            res[1] = r+1
            return res
        }
    }

    return res
}
