type Car struct {
        time float64
        pos int
}

func carFleet(target int, position []int, speed []int) int {
    

    var res int

    cars := make([]Car, len(speed))
    for i:=0; i < len(speed); i++ {
        cars[i] = Car {time : (float64(target) - float64(position[i]))/float64(speed[i]), pos: position[i]}
    }

    sort.Slice(cars, func(a, b int) bool {
        return cars[a].pos > cars[b].pos
    })

    fmt.Println(cars)
    
    var max float64
    for _,c := range cars {
        if c.time > max {
            res++
            max = c.time
            fmt.Println("increment res on p: ", c)
        }
    }

    return res
}
