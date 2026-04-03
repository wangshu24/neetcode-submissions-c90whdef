func combinationSum2(cand []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(cand)
	var back func([]int, int, int)
	back = func(set []int, start, tar int) {
		if tar == target {
			add := make([]int,len(set))
			copy(add, set)
			res = append(res, add)
			return
		} else if tar > target || start >= len(cand) {return}
		for i:=start; i < len(cand); i++ {
			if i>start && cand[i] == cand[i-1] {
				continue
			}
			set = append(set, cand[i])
			back(set, i+1, tar+cand[i])
			set = set[:len(set)-1]
		}
	}
	back([]int{}, 0, 0)
	return res
}
