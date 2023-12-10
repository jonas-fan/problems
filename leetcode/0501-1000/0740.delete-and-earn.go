// #dynamic-programming

import (
    "sort"
)

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func counts(nums []int) map[int]int {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num] += num
    }

    return seen
}

func deleteAndEarn(nums []int) int {
    points := counts(nums)

    nums = nums[:0]

    for num := range points {
        nums = append(nums, num)
    }

    if len(nums) == 1 {
        return points[nums[0]]
    }

    sort.Ints(nums)

    dp := make([]int, len(nums))

    dp[0] = points[nums[0]]

    if nums[1] - nums[0] > 1 {
        dp[1] = dp[0] + points[nums[1]]
    } else {
        dp[1] = max(dp[0], points[nums[1]])
    }

    for i := 2; i < len(nums); i++ {
        if nums[i] - nums[i-1] > 1 {
            dp[i] = dp[i-1] + points[nums[i]]
        } else {
            dp[i] = max(dp[i-1], dp[i-2] + points[nums[i]])
        }
    }

    return dp[len(nums)-1]
}
