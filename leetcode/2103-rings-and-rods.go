// #string

import (
    "math/bits"
)

func countPoints(rings string) int {
    rods := [10]uint{}

    for i := 0; i < len(rings); i +=2 {
        what := rings[i] - 'A'
        where := rings[i+1] - '0'

        rods[where] |= uint(1) << what
    }

    out := 0

    for _, count := range rods {
        if bits.OnesCount(count & ^uint(0)) > 2 {
            out++
        }
    }

    return out
}
