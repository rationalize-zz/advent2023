package main

import (
    "fmt"
    "os"
    "strings"
    "image"
)

type pipe struct {
    exit1 image.Point
    exit2 image.Point
}

var movemap = map[byte]pipe {
    '|': pipe { image.Point {0, 1}, image.Point {0, -1} },
    '-': pipe { image.Point {1, 0}, image.Point {-1, 0} },
    'L': pipe { image.Point {0, -1}, image.Point {1, 0} },
    'J': pipe { image.Point {0, -1}, image.Point {-1,0} },
    '7': pipe { image.Point {0, 1}, image.Point {-1, 0} },
    'F': pipe { image.Point {0, 1}, image.Point {1, 0} },
}

func moveLoc(loc image.Point, mod image.Point) image.Point {
    return image.Point { loc.X + mod.X, loc.Y + mod.Y }
}

func diffLoc(l1 image.Point, l2 image.Point) bool {
    if l1.X == l2.X && l1.Y == l2.Y {
        return false
    }
    return true
}

func findStartLoc(data []string) image.Point {
    for i, line := range data {
        for j := 0; j < len(line); j++ {
            if line[j] == 'S' {
                return image.Point {j, i}
            }
        }
    }
    panic("Bad dataset")
}

func isBackflow(s image.Point, n image.Point, p pipe) bool {
    if !diffLoc(s, moveLoc(n, p.exit1)) || !diffLoc(s, moveLoc(n, p.exit2)) {
        return true
    }
    return false
}

func getStartPipeSym(s image.Point, data []string) byte {
    for key, pipe := range movemap {
        n1 := moveLoc(s, pipe.exit1)
        n2 := moveLoc(s, pipe.exit2)
        if isBackflow(s, n1, movemap[data[n1.Y][n1.X]]) && isBackflow(s, n2, movemap[data[n2.Y][n2.X]]) {
            return key
        }
    }
    panic("nopipematch")
}

func findfill(up byte, down byte, left byte, right byte) byte {
    if up == '7' || up == '|' || up == 'F' {
        if down == 'J' || down == '|' || down == 'L' {
            return '|'
        }
    }
    if left == 'F' || left == '-' || left == 'L' {
        if right == '7' || right == '-' || right == 'J' {
            return '-'
        }
    }
    return ','
}

func main() {
    
    input, _ := os.ReadFile("./input")
    lines := strings.Fields(string(input))
    bytearray := make([][]byte, 0)
    ph := make([]byte, (len(lines[0])*2) + 1)
    for i := 0; i < len(ph); i++ {
        ph[i] = ','
    }
    bytearray = append(bytearray, ph)
    for _, line := range lines {
        byteline := make([]byte, (len(line)*2) + 1)
        for i := 0; i < len(line); i++ {
            byteline[i*2] = ','
            byteline[i*2+1] = '.'
        }
        byteline[len(byteline)-1] = ','
        bytearray = append(bytearray, byteline)
        ph := make([]byte, (len(lines[0])*2) + 1)
        for i := 0; i < len(ph); i++ {
            ph[i] = ','
        }
        bytearray = append(bytearray, ph)
    }
    stepcount := 1
    s := findStartLoc(lines)
    p := s
    startpipesym := getStartPipeSym(s, lines)
    bytearray[s.Y*2+1][s.X*2+1] = startpipesym
    c := moveLoc(s, movemap[startpipesym].exit1)
    for diffLoc(c, s) {
        bytearray[c.Y*2+1][c.X*2+1] = lines[c.Y][c.X]
        pipe := movemap[lines[c.Y][c.X]]
        n := moveLoc(c, pipe.exit1)
        if !diffLoc(n, p) {
            n = moveLoc(c, pipe.exit2)
        }
        p, c = c, n
        stepcount++
    }
    stepcount = stepcount >> 1
    fmt.Println("P1:", stepcount)
    for i := 1; i < len(bytearray)-1; i++ {
        for j := 1; j < len(bytearray[i]) - 1; j++ {
            if bytearray[i][j] == ',' {
                bytearray[i][j] = findfill(bytearray[i-1][j], bytearray[i+1][j], bytearray[i][j-1], bytearray[i][j+1])
            }
        }
    }
    ymax := len(bytearray)
    xmax := len(bytearray[0])
    fillstack := []image.Point{{0,0}}
    for len(fillstack) > 0 {
        c := fillstack[len(fillstack) - 1]
        fillstack = fillstack[:len(fillstack) - 1]
        if bytearray[c.Y][c.X] != '.' && bytearray[c.Y][c.X] != ',' {
            continue
        }
        bytearray[c.Y][c.X] = '0'
        if c.Y - 1 >= 0 {
            fillstack = append(fillstack, image.Point {c.X, c.Y - 1})
        }
        if c.Y + 1 < ymax {
            fillstack = append(fillstack, image.Point {c.X, c.Y + 1})
        }
        if c.X - 1 >= 0 {
            fillstack = append(fillstack, image.Point {c.X - 1, c.Y})
        }
        if c.X + 1 < xmax {
            fillstack = append(fillstack, image.Point {c.X + 1, c.Y})
        }
    }
    dotcount := 0;
    for _, row := range bytearray {
        for _, col := range row {
            if col == '.' {
                dotcount++
            }
        }
    }
    fmt.Println("P2:", dotcount)
}
