package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

type maprange struct {
    srcstart uint64
    srcend uint64
    deststart uint64
}

type mapping struct {
    source string
    dest string
    ranges []maprange
}

func getMapRange(data []string) maprange {
    first, _ := strconv.ParseUint(data[0], 10, 64)
    second, _ := strconv.ParseUint(data[1], 10, 64)
    third, _ := strconv.ParseUint(data[2], 10, 64)
    return maprange { srcstart: second, srcend: second + third - 1, deststart: first }
}

func traverseXform(xform *mapping, val uint64) (string, uint64) {
    ret := val
    for _, mrange := range xform.ranges {
        if val >= mrange.srcstart && val <= mrange.srcend {
            ret = val - mrange.srcstart + mrange.deststart
            break
        }
    }
    return xform.dest, ret
}

func main() {

    input, _ := os.ReadFile("./input")
    textlines := strings.Split(string(input), "\n")
    seedlist := make([]uint64, 0)
    var xforms = make(map[string]*mapping)
    _, textseeds, _ := strings.Cut(textlines[0], ": ")

    seeds := strings.Split(textseeds, " ")
    for _, seed := range seeds {
        temp, _ := strconv.ParseUint(seed, 10, 64)
        seedlist = append(seedlist, temp)
    }

    needheader := true
    src := ""
    for _, line := range textlines {
        if needheader {
            info, test, valid := strings.Cut(line, " ")
            if !valid || test != "map:" {
                continue
            }
            headerdata := strings.Split(info, "-")
            src = headerdata[0]
            xforms[headerdata[0]] = &mapping { source: headerdata[0], dest: headerdata[2], ranges: make([]maprange, 0) }
            needheader = false
        } else {
            data := strings.Split(line, " ")
            if len(data) != 3 {
                needheader = true
                continue
            }
            xforms[src].ranges = append(xforms[src].ranges, getMapRange(data))
        }
    }
    var lowest_p1 uint64 = math.MaxUint64
    for _, seed := range seedlist {
        xformstring := "seed"
        current := seed
        for xformstring != "location" {
            xformstring, current = traverseXform(xforms[xformstring], current)
        }
        if current < lowest_p1 {
            lowest_p1 = current
        }
    }
    var lowest_p2 uint64 = math.MaxUint64
    for i := 0; i < len(seedlist); i+=2 {
        fmt.Println("Origin ", i, " seed: ", seedlist[i], " range: ", seedlist[i+1])
        for j := seedlist[i]; j < seedlist[i] + seedlist[i+1] - 1; j++ {
            xformstring := "seed"
            current := j
            for xformstring != "location" {
                xformstring, current = traverseXform(xforms[xformstring], current)
            }
            if current < lowest_p2 {
                lowest_p2 = current
            }
        }
    }
    fmt.Println("P1: ", lowest_p1, "\nP2: ", lowest_p2)
}
