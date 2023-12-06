package main

import (
    "fmt"
    "os"
    "strings"
    "math"
)

func getvals(numstr string) []int {
    active, value := false, 0
    vals := make([]int, 0)
    for i := 0; i < len(numstr); i++ {
        if numstr[i] >= '0' && numstr[i] <= '9' {
            if !active {
                active = true
            }
            value = (value * 10) + (int)(numstr[i] - '0')
        } else {
            if active {
                vals = append(vals, value)
                active, value = false, 0
            }
        }
    }
    if active {
        vals = append(vals, value)
    }
    return vals
}

func getval_p2(numstr string) int {
    value := 0
    for i := 0; i < len(numstr); i++ {
        if numstr[i] >= '0' && numstr[i] <= '9' {
            value = (value * 10) + (int)(numstr[i] - '0')
        }
    }
    return value
}

func getLimits(time int, distance int) (int, int) {
    var bulk = math.Sqrt(math.Pow(float64(time), 2.0) - 4.0*float64(distance))
    var low = (float64(time) - bulk) / 2.0
    var high = (float64(time) + bulk) / 2.0
    lowret := math.Ceil(low)
    highret := math.Floor(high)
    if low == lowret { //meets distance, but doesn't BEAT distance
        lowret = lowret + 1.0
    }
    if high == highret {
        highret = highret - 1.0
    }
    return int(lowret), int(highret)
}

func main() {
    
    collect_p1 := 1
    input, _ := os.ReadFile("./input")
    textlines := strings.Split(string(input), "\n")
    times := getvals(textlines[0])
    distances := getvals(textlines[1])
    for i := 0; i < len(times); i++ {
        low, high := getLimits(times[i], distances[i])
        collect_p1 *= (high - low + 1)
    }
    low2, high2 := getLimits(getval_p2(textlines[0]), getval_p2(textlines[1]))

    fmt.Println("P1: ", collect_p1, "\nP2: ", high2 - low2 + 1)
}
