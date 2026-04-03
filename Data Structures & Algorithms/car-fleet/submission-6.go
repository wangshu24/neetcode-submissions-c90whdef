
func carFleet(target int, position []int, speed []int) int {
    type car struct {
        pos int
        time float64
    }
    
    cars := make([]car, len(speed)) 
    for i:=0; i < len(speed); i++ {
        cars[i] = car{ position[i], float64(target - position[i]) / float64(speed[i])}
    }    
    fmt.Println(cars)
    sort.Slice(cars, func(i, j int) bool {
        return cars[i].pos > cars[j].pos
    })

    fleet, max := 0, 0.0    
    for _,c := range cars {
        if c.time > max {
            fleet++
            max = c.time
        }
    }
    // fmt.Println(time)
    fmt.Println(cars)
    return fleet
}
