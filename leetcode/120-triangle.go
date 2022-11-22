// #dynamic-programming

import (
    "math"
)

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func minimumTotal(triangle [][]int) int {
    sum := make([]int, len(triangle[len(triangle)-1]))
    sum[0] = triangle[0][0]

    for row := 0; row < len(triangle) - 1; row++ {
        for col := len(triangle[row]); col >= 0; col-- {
            val := triangle[row+1][col]

            if col == 0 {
                sum[col] = sum[col] + val
            } else if col == len(triangle[row]) {
                sum[col] = sum[col-1] + val
            } else {
                sum[col] = min(sum[col]+val, sum[col-1]+val)
            }
        }
    }

    out := math.MaxInt

    for _, val := range sum {
        out = min(out, val)
    }

    return out
}
