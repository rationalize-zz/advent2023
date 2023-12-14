package main

import (
    "fmt"
    "os"
    "strings"
    "image"
    "math"
)

func main() {
    input, _ := os.ReadFile("./input")
    textlines := strings.Fields(string(input))
    fmt.Println("P1:", getvals(2, textlines), "\nP2:", getvals(1000000, textlines))
}

func getvals(expand_rate int, data []string) int {
    collect, rinc, cinc := 0, 0, 0
    rmap := make([]int, len(data)) 
    cmap := make([]int, len(data[0]))
    graw := []image.Point{}

    for i, line := range data {
        rtest := false
        for j := 0; j < len(line); j++ {
            if line[j] == '#' {
                rtest = true
                cmap[j] = 1
                graw = append(graw, image.Point{j, i})
            }
        }
        if !rtest {
            rinc += (expand_rate - 1)
        }
        rmap[i] = rinc + i
    }

    for i, c := range cmap {
        if c == 0 {
            cinc += (expand_rate - 1)
        }
        cmap[i] = cinc + i
    }

    for len(graw) > 1 {
        s := graw[len(graw) - 1]
        graw = graw[:len(graw) - 1]
        for _, g := range graw {
            collect += int(math.Abs(float64(cmap[s.X] - cmap[g.X])) + math.Abs(float64(rmap[s.Y] - rmap[g.Y])))
        }
    }
    return collect
}
