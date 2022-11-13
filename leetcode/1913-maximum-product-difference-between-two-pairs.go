// #array

import (
    "sort"
)

func maxProductDifference(nums []int) int {
    sort.Ints(nums)

    min := nums[0] * nums[1]
    max := nums[len(nums)-1] * nums[len(nums)-2]

    return max - min
}
