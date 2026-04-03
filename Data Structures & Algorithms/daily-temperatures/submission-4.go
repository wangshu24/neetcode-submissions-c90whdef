func dailyTemperatures(temps []int) []int {
    res, stack := make([]int,len(temps)), []int{}

    for i, temp := range temps {
        for len(stack) > 0 && temp > temps[stack[len(stack)-1]] {
            stackInd := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[stackInd] = i - stackInd
        }
        stack = append(stack, i)
    }
    return res
}
