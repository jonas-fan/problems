// #dynamic-programming

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func minCostClimbingStairs(cost []int) int {
    dp := make([]int, 0, len(cost)+1)

    dp = append(dp, 0)
    dp = append(dp, 0)

    for i := 2; i <= len(cost); i++ {
        dp = append(dp, min(dp[i-1] + cost[i-1],
                            dp[i-2] + cost[i-2]))
    }

    return dp[len(dp)-1]
}
