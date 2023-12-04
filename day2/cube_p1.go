package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func hand_is_valid (hand string) bool {
    counts := strings.Split(hand, ", ")
    for _, count := range counts {
        val, color, _ := strings.Cut(count, " ")
        temp, err := strconv.Atoi(val)
        if err != nil {
            panic(err)
        }
        switch len(color) {
        case 3: //red
            if temp > 12 {
                return false
            }
        case 4:
            if temp > 14 {
                return false
            }
        case 5:
            if temp > 13 {
                return false
            }
        default:
            panic("Error: fix your color parsing")
        }
    }
    return true
}

func main() {
    
    var collector int = 0

    input, err := os.ReadFile("../input")
    if err != nil {
        panic(err)
    }

    gamelist := strings.Split(string(input), "\n")

    for _, game := range gamelist {
        gameinfo, handlist, valid := strings.Cut(game, ": ")
        if !valid {
            continue
        }
        hands := strings.Split(handlist, "; ")
        validgame := true
        for _, hand := range hands {
            if !hand_is_valid(hand) {
                validgame = false
                break
            }
        }
        if !validgame {
            continue
        }
        _, gamenum, _ := strings.Cut(gameinfo, " ")
        temp, err := strconv.Atoi(gamenum)
        if err != nil {
            panic(err)
        }
        collector += temp
    }
    fmt.Println("GameVal: ", collector)
}


