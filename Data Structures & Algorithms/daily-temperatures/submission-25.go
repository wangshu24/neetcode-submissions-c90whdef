func dailyTemperatures(temp []int) []int {
    res, days := make([]int, len(temp)), make([]int,0)

    for i,t := range temp {
        for len(days) > 0 && t > temp[days[len(days)-1]] {
            res[days[len(days)-1]] = i - days[len(days)-1]
            days = days[:len(days)-1]
        }
        days = append(days, i)
    }

    return res
}
