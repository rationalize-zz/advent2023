package main

import (
    "fmt"
    "os"
    "strings"
)

func diffCount(l string, r string) int {
    count:=0
    for i:=0; i < len(l); i++ {
        if l[i] != r[i] {
            count++
        }
    }
    return count
}

func hscan(field []string, smudge int) (bool, int) {
    for i:=0; i<len(field)-1; i++{
        l, r, e:=i, i+1, 0
        for e <= smudge {
            if l < 0 || r == len(field) {
                if e == smudge {
                    return true, i+1
                }
                break
            }
            e += diffCount(field[l], field[r])
            l, r=l-1, r+1
        }
    }
    return false, 0
}

func vscan(field []string, smudge int) (bool,int) {
    tilted := make([]string, len(field[0]))
    for _, row := range field {
        for i, pos := range row {
            tilted[i] += string(pos)
        }
    }
    return hscan(tilted, smudge)
}

func main() {
    input, _ := os.ReadFile("./input")
    lavafields := strings.Split(string(input), "\n\n")
    fmt.Println("P1:", solve(lavafields, 0))
    fmt.Println("P2:", solve(lavafields, 1))
}

func solve(lavafields []string, smudge int) int {
    collect:=0
    for _, blobfield := range lavafields { //for each lava field
        lfield := strings.Fields(blobfield)
        valid, temp := hscan(lfield, smudge)
        if valid {
            collect += (temp*100)
            continue
        }
        valid, temp = vscan(lfield, smudge)
        if (valid) {
            collect += temp
            continue
        }
        panic("Should never get here")
    }
    return collect
}
