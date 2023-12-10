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
    
    var collect_p2 uint64 = 1
    nodemap := make(map[string]nodeinfo)
    startlocs := make([]string, 0)
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
        if loc[2] == 'A' {
            startlocs = append(startlocs, loc)
        }
        nodemap[loc] = nodeinfo { left: left, right: right}
    }
    
    steplen := len(stepseq)

    for i := 0; i < len(startlocs); i++ {
        stepind := 0
        var steps uint64 = 0
        for startlocs[i][2] != 'Z' {
            if stepseq[stepind] == 'R' {
                startlocs[i] = nodemap[startlocs[i]].right
            } else {
                startlocs[i] = nodemap[startlocs[i]].left
            }
            stepind++
            steps++
            if stepind == steplen {
                stepind = 0
            }
        }
        steps /= uint64(steplen)
        collect_p2 *= steps
    }
    collect_p2 *= uint64(steplen)
    fmt.Println("P2:", collect_p2)
}
