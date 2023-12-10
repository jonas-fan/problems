// #array

import (
    "sort"
)

func sortedSquares(nums []int) []int {
    out := make([]int, 0, len(nums))

    for _, num := range nums {
        out = append(out, num * num)
    }

    sort.Ints(out)

    return out
}
