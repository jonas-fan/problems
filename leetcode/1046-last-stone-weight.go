// #array

import (
    "sort"
)

func lastStoneWeight(stones []int) int {
    for len(stones) > 1 {
        sort.Ints(stones)

        size := len(stones)

        stones[size-2] = stones[size-1] - stones[size-2]

        if stones[size-2] == 0 {
            stones = stones[:size-2]
        } else {
            stones = stones[:size-1]
        }
    }

    if len(stones) == 0 {
        return 0
    }

    return stones[0]
}
