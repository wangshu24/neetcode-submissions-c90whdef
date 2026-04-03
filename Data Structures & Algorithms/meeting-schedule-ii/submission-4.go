
/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {return 0}
	
								
	starts, ends := []int{}, []int{}
	for _, interval := range intervals {
		starts = append(starts, interval.start)
		ends = append(ends, interval.end)
	}
	sort.Ints(starts)
	sort.Ints(ends)
	fmt.Println(starts, ends)

	rooms := 0
	s_ptr, e_ptr := 0,0
	for s_ptr < len(starts) && e_ptr < len(ends) {
		if starts[s_ptr] < ends[e_ptr] {	
			rooms++
		} else {
			e_ptr++
		} 
		s_ptr++
	}
	return rooms
}

