package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

func main() {

    var collector int = 0

    parse := regexp.MustCompile("[0-9]")

    input, err := os.ReadFile("./input")
    if err != nil {
        panic(err)
    }
    
    tokens := strings.Split(string(input), "\n")

    for _, line := range tokens {
        nums := parse.FindAllString(line, -1)
        if len(nums) == 0 {
            continue
        }
        val, err := strconv.Atoi(nums[0])
        if err != nil {
            panic(err)
        }
        collector += val * 10
        val, err = strconv.Atoi(nums[len(nums)-1])
        if err != nil {
            panic(err)
        }
        collector += val
    }
    fmt.Println("Calibration: ", collector)
}
