// #string

import (
    "math"
)

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func shortestToChar(s string, c byte) []int {
    out := make([]int, len(s))
    idx := make([]int, 0, len(s))

    for i, letter := range s {
        if letter == rune(c) {
            out[i] = 0
            idx = append(idx, i)
        } else {
            out[i] = math.MaxInt
        }
    }

    for _, i := range idx {
        left, right := i - 1, i + 1

        for left >= 0 && out[left] != 0 {
            out[left] = min(out[left], i-left)
            left--
        }

        for right < len(out) && out[right] != 0 {
            out[right] = min(out[right], right-i)
            right++
        }
    }

    return out
}
