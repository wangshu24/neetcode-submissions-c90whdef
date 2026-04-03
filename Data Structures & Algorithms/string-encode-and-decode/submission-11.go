type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    var res string
    for _,i := range strs {
        res += strconv.Itoa(len(i)) + "#" + i 
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    if len(encoded) == 0 {
        return []string{}
    }
    i:=0
    res := []string{}
    for i < len(encoded) {
        j := i 
        for encoded[j] != '#' {
            j++
        }
        length,_ := strconv.Atoi(encoded[i:j])
        i = j+1
        res = append(res, encoded[i: i+length])
        i= i+length
    }

    return res
}
