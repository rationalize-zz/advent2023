package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func getNextVal(seq []int) int {
    nextseq := make([]int, 0)
    seqlen := len(seq)
    allzeros := true
    for i := 1; i < seqlen; i++ {
        val := seq[i] - seq[i - 1]
        if val != 0 {
            allzeros = false
        }
        nextseq = append(nextseq, val)
    }
    if allzeros {
        return seq[seqlen-1]
    }
    return seq[seqlen-1] + getNextVal(nextseq)
}

func main() {
    
    collect_p1, collect_p2 := 0, 0
    input, _ := os.ReadFile("./input")
    lines := strings.Split(string(input), "\n")
    for _, line := range lines {
        lineseq := make([]int, 0)
        textvals := strings.Split(line, " ")
        if len(textvals) == 0 {
            continue
        }
        for _, tval := range textvals {
            val, _ := strconv.Atoi(tval)
            lineseq = append(lineseq, val)
        }
        collect_p1 += getNextVal(lineseq)
        for i, j := 0, len(lineseq) - 1; i < j; i, j = i+1, j-1 {
            lineseq[i], lineseq[j] = lineseq[j], lineseq[i]
        }
        collect_p2 += getNextVal(lineseq)
    }

    fmt.Println("P1:", collect_p1, "\nP2", collect_p2)
}
