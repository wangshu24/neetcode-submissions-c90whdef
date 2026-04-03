func carFleet(target int, position []int, speed []int) int {
    type car struct {
        pos  int
        time float64
    }
    cars := make([]car, len(position))
    for i := 0; i < len(position); i++ {
        cars[i] = car{position[i], float64(target-position[i]) / float64(speed[i])}
    }
    sort.Slice(cars, func(i, j int) bool {
        return cars[i].pos > cars[j].pos
    })

    fleets := 0
    maxTime := 0.0
    for _, c := range cars {
        if c.time > maxTime {
            fleets++
            maxTime = c.time
        }
    }
    return fleets
}