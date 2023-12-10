// #array

import (
    "sort"
)

func sorts(nums []int) []int {
    sorted := make([]int, len(nums))

    copy(sorted, nums)

    sort.Ints(sorted)

    return sorted
}

func heightChecker(heights []int) int {
    expected := sorts(heights)
    diff := 0

    for i := range heights {
        if heights[i] != expected[i] {
            diff++
        }
    }

    return diff
}
