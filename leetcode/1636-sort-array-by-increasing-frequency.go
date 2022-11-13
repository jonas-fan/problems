// #array

import (
    "sort"
)

func frequencySort(nums []int) []int {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num]++
    }

    sort.Slice(nums, func(lhs int, rhs int) bool {
        if seen[nums[lhs]] != seen[nums[rhs]] {
            return seen[nums[lhs]] < seen[nums[rhs]]
        }

        return nums[lhs] > nums[rhs]
    })

    return nums
}
