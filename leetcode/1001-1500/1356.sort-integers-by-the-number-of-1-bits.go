// #array

import (
    "math/bits"
    "sort"
)

func sortByBits(arr []int) []int {
    sort.Slice(arr, func(lhs int, rhs int) bool {
        lones := bits.OnesCount(uint(arr[lhs]))
        rones := bits.OnesCount(uint(arr[rhs]))

        if lones == rones {
            return arr[lhs] < arr[rhs]
        }

        return lones < rones
    })

    return arr
}
