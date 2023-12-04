package main

import (
    "fmt"
    "os"
    "strings"
)

type partnumber struct {
    value int
    index int
    length int
}

type linedata struct {
    parts []partnumber
    syms []int
    gears []int
}

func getlinedata(line string) linedata {
    active, value, index, vallen := false, 0, 0, 0
    ld := linedata { parts: make([]partnumber, 0), syms: make([]int, 0), gears: make([]int, 0) }
    for i := 0; i < len(line); i++ {
        if line[i] >= '0' && line[i] <= '9' {
            if !active {
                active, index = true, i
            }
            value = (value * 10) + (int)(line[i] - '0')
            vallen++
        } else {
            if active {
                ld.parts = append(ld.parts, partnumber {value: value, index: index, length: vallen })
                active, value, vallen = false, 0, 0
            }
            if line[i] == '*' {
                ld.gears = append(ld.gears, i)
            } else if line[i] != '.' {
                ld.syms = append(ld.syms, i)
            } 
        }
    }
    if active {
        ld.parts = append(ld.parts, partnumber {value: value, index: index, length: vallen })
    }
    return ld
}

func parseds_p1 (ds []linedata) int {
    collect := 0
    symlist := make([]int, 0)
    for _, ld := range ds {
        symlist = append(append(symlist, ld.syms...), ld.gears...)
    }
    for _, part := range ds[len(ds)-2].parts {
        for _, sym := range symlist {
            if part.index <= sym + 1 && part.index + part.length - 1 >= sym - 1 {
                collect += part.value
                break
            }
        }
    }
    return collect
}

func parseds_p2 (ds []linedata) int {
    collect := 0
    partlist := make([]partnumber, 0)
    for _, ld := range ds {
        partlist = append(partlist, ld.parts...)
    }
    for _, sym := range ds[len(ds)-2].gears {
        matches, ratio := 0, 1
        for _, part := range partlist {
            if part.index <= sym + 1 && part.index + part.length - 1 >= sym - 1 {
                matches++
                ratio *= part.value
            }
        }
        if matches == 2 {
            collect += ratio
        }
    }
    return collect
}

func main() {
    collect_p1, collect_p2 := 0, 0
    input, _ := os.ReadFile("./input")
    datastore := make([]linedata, 0)
    linelist := strings.Split(string(input), "\n")

    for _, line := range linelist {
        datastore = append(datastore, getlinedata(line))
        if len(datastore) > 3 {
            datastore = datastore[1:]
        } else if len(datastore) < 2 {
            continue
        }
        collect_p1 += parseds_p1(datastore)
        collect_p2 += parseds_p2(datastore)
    }
    fmt.Println("P1: ", collect_p1, "\nP2: ", collect_p2)
}
