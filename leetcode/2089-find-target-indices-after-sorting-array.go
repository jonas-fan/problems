// #array

import (
    "sort"
)

func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)

    out := []int{}

    for i, num := range nums {
        if num == target {
            out = append(out, i)
        }
    }

    return out
}
