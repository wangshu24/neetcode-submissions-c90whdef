func dailyTemperatures(temperatures []int) []int {
    var res []int = make([]int, len(temperatures))
    var days []int = make([]int, 0)
    for d, t := range temperatures { 
            for len(days) > 0 && t > temperatures[days[len(days)-1]] {
                res[days[len(days)-1]] = d - days[len(days)-1]
                days = days[:len(days)-1]
            }
            days = append(days,d)
    }
    return res
}
