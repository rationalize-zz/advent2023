package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    
    var collector int = 0

    input, err := os.ReadFile("../input")
    if err != nil {
        panic(err)
    }

    gamelist := strings.Split(string(input), "\n")

    for _, game := range gamelist {
        red := 0
        blue := 0
        green := 0
        _, handlist, valid := strings.Cut(game, ": ")
        if !valid { //cuts off blank lines
            continue
        }
        hands := strings.Split(handlist, "; ")
        for _, hand := range hands {
            counts := strings.Split(hand, ", ")
            for _, count := range counts {
                tval, color, _ := strings.Cut(count, " ")
                val, err := strconv.Atoi(tval)
                if err != nil {
                    panic(err)
                }
                switch len(color) {
                case 3:
                    if val > red {
                        red = val
                    }
                case 4:
                    if val > blue {
                        blue = val
                    }
                case 5:
                    if val > green {
                        green = val
                    }
                default:
                    panic("Error: parsing went bad")
                }
            }
        }

        collector += (red * blue * green)
    }
    fmt.Println("GameVal: ", collector)
}
