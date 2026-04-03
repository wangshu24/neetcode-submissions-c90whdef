type TimeMap struct {
	keys map[string][]pair
}

type pair struct {
	time int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		keys: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keys[key] = append(this.keys[key], pair{time: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	
	if _, there := this.keys[key]; !there {
		return ""
	}

	pair:= this.keys[key]
	l, r := 0, len(pair)-1
	for l <= r {
		mid := l + (r-l)/2
		
		if pair[mid].time <= timestamp {
			if mid == len(pair)-1 || pair[mid+1].time > timestamp {
				return pair[mid].value
			}
			l = mid + 1
		}  else {
			r = mid - 1
		}
	}
	
	return ""
}
