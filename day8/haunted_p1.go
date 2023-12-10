package main

import (
    "fmt"
    "os"
    "strings"
)

type nodeinfo struct {
    left string
    right string
}

func main() {
    
    collect_p1 := 0
    nodemap := make(map[string]nodeinfo)
    input, _ := os.ReadFile("./input")
    stepseq, nodeblock, _ := strings.Cut(string(input), "\n\n")
    nodelines := strings.Split(nodeblock, "\n")

    for _, nodetext := range nodelines {
        loc, destinations, v := strings.Cut(nodetext, " = (")
        if !v {
            continue
        }
        destinations, _ = strings.CutSuffix(destinations, ")")
        left, right, _ := strings.Cut(destinations, ", ")
        nodemap[loc] = nodeinfo { left: left, right: right}
    }
    
    loc := "AAA"
    stepind := 0
    steplen := len(stepseq)

    for loc != "ZZZ" {
        if stepseq[stepind] == 'R' {
            loc = nodemap[loc].right
        } else {
            loc = nodemap[loc].left
        }
        stepind++
        if stepind == steplen {
            stepind = 0
        }
        collect_p1++
    }

    fmt.Println("P1:", collect_p1)
}
