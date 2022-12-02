// #math

import (
    "math"
)

func judgeSquareSum(c int) bool {
    if c == 0 {
        return true
    }

    left := 0
    right := int(math.Sqrt(float64(c)))

    for left <= right {
        lval := left * left
        rval := right * right

        if lval + rval == c {
            return true
        } else if lval + rval < c {
            left++
        } else {
            right--
        }
    }

    return false
}
