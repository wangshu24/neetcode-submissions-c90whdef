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
	
	idx := sort.Search(len(pair), func(i int ) bool {
		return pair[i].time > timestamp
	})

	if idx == 0 {
		return  ""
	}

	return pair[idx-1].value
}
