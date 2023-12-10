// #array

import (
    "sort"
)

func sum(nums []int) int {
    out := 0

    for _, num := range nums {
        out += num
    }

    return out
}

func minSubsequence(nums []int) []int {
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    total := sum(nums)
    subtotal := 0

    for i := 0; i < len(nums); i++ {
        subtotal += nums[i]

        if subtotal > (total - subtotal) {
            return nums[:i+1]
        }
    }

    return []int{}
}
