import ("slices")
/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	starts, ends := make([]int,0), make([]int,0)

	for _, inter := range intervals {
		starts = append(starts, inter.start)
		ends = append(ends, inter.end)
	}
	slices.Sort(starts)
	slices.Sort(ends)
	fmt.Println(starts, ends)

	var room int
	sptr, eptr := 0, 0
	for sptr < len(intervals) && eptr < len(intervals) {
		if ends[eptr] > starts[sptr] {
			room++
		} else {
			eptr++
		}
			sptr++
	}


	return room
}
