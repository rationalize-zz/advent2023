package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

type handinfo struct {
    rank int
    val int
    bid int
}

func cmp_handinfo(a, b handinfo) int {
    if a.rank != b.rank {
        return a.rank - b.rank
    }
    return a.val - b.val
}

func parsehand(t_bid string, t_hand string, p2 bool) handinfo {
    bid, _ := strconv.Atoi(t_bid)
    val, rank, maxcount, jokers := 0, 0, 0, 0
    handtest := make(map[byte]int)
    for i := 0; i < len(t_hand); i++ {
        val = val * 100
        switch t_hand[i] {
        case 'T':
            val += 10
        case 'J':
            if p2 {
                val += 1
                jokers += 1
                continue
            } else {
                val += 11
            }
        case 'Q':
            val += 12
        case 'K':
            val += 13
        case 'A':
            val += 14
        default:
            val += int(t_hand[i] - '0')
        }
        handtest[t_hand[i]] += 1
        if handtest[t_hand[i]] > maxcount {
            maxcount = handtest[t_hand[i]]
        }
    }
    maxcount += jokers
    switch len(handtest) {
    case 0:
        rank = 6 //all jokers, 5 of a kind
    case 1:
        rank = 6
    case 2:
        if maxcount == 4 {
            rank = 5
        } else { 
            rank = 4
        }
    case 3:
        if maxcount == 3 {
            rank = 3
        } else {
            rank = 2
        }
    case 4:
        rank = 1
    case 5:
        rank = 0
    default:
        panic("How?!?!?")
    }
    return handinfo { rank: rank, val: val, bid: bid }
}

func main() {
    
    collect_p1, collect_p2 := 0, 0
    input, _ := os.ReadFile("./input")
    textlines := strings.Split(string(input), "\n")
    hands_p1 := make([]handinfo, 0)
    hands_p2 := make([]handinfo, 0)
    for _, line := range textlines {
        texthand, textbid, valid := strings.Cut(line, " ")
        if !valid {
            continue
        }
        hands_p1 = append(hands_p1, parsehand(textbid, texthand, false))
        hands_p2 = append(hands_p2, parsehand(textbid, texthand, true))
    }

    slices.SortFunc(hands_p1, cmp_handinfo)
    slices.SortFunc(hands_p2, cmp_handinfo)

    for i, hand := range hands_p1 {
        collect_p1 += hand.bid * (i + 1)
    }
    for i, hand := range hands_p2 {
        collect_p2 += hand.bid * (i + 1)
    }

    fmt.Println("P1:", collect_p1, "\nP2:", collect_p2)
}
