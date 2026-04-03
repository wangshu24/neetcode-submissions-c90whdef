type Area struct {
	pref int
	suf int
}

func largestRectangleArea(heights []int) int {
	res := math.MinInt32
	areas := make([]Area, 0)
	for i:=0; i < len(heights); i++ {
		l,r := 0, 0
		for i - l-1 >= 0 && heights[i] <= heights[i-l-1] {
			l++
		}
		for i + r+1 < len(heights) && heights[i] <= heights[i+r+1] {
			r++
		}
		areas = append(areas, Area{pref: l, suf: r})
	}
	fmt.Println(areas)

	for i:=0; i  < len(heights); i++ {
		res = max(res, heights[i] * (areas[i].pref + areas[i].suf + 1))
	}

	return res
}
