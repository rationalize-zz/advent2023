package main

import (
    "fmt"
    "os"
    "strings"
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

func main() {

    collect_p1, collect_p2 := 0, 0

    input, err := os.ReadFile("./input")
    if err != nil {
        panic(err)
    }
    
    multi := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    cards := strings.Split(string(input), "\n")

    for _, card := range cards {
        _, data, valid := strings.Cut(card, ": ")
        if !valid {
            continue
        }
        wins, tries, _ := strings.Cut(data, " | ")
        winvals := getvals(wins)
        tryvals := getvals(tries)
        cardmatches := 0
        for _, tryval := range tryvals {
            for _, winval := range winvals {
                if tryval == winval {
                    cardmatches++
                }
            }
        }
        if cardmatches != 0 {
            points := 1 << (cardmatches - 1)
            collect_p1 += points
        }
        multiply := 1 + multi[0] 
        multi = multi[1:]
        multi = append(multi, 0)

        collect_p2 += multiply

        for i := 0; i < cardmatches; i++ {
            multi[i] = multi[i] + multiply
        }
    }
    fmt.Println("P1: ", collect_p1, "\nP2: ", collect_p2)
}

