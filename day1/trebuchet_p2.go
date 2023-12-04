package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    var collector int = 0

    nummap := map[string]int {
        "1": 1,
        "one": 1,
        "2": 2,
        "two": 2,
        "3": 3,
        "three": 3,
        "4": 4,
        "four": 4,
        "5": 5,
        "five": 5,
        "6": 6,
        "six": 6,
        "7": 7,
        "seven": 7,
        "8": 8,
        "eight": 8,
        "9": 9,
        "nine": 9,
    }

    input, err := os.ReadFile("./input")
    if err != nil {
        panic(err)
    }
    
    tokens := strings.Split(string(input), "\n")

    for _, line := range tokens {
        var first = 0
        var firstind = 10000
        var last = 0
        var lastind = -1

        for text, number := range nummap {
            temp := strings.Index(line, text)
            if temp != -1 && temp < firstind {
                firstind = temp
                first = number
            }
            temp = strings.LastIndex(line, text)
            if temp != -1 && temp > lastind {
                lastind = temp
                last = number
            }
        }

        if first == 0 || last == 0 {
            continue
        }

        collector += first * 10
        collector += last
    }
    fmt.Println("Calibration: ", collector)
}
