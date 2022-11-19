// #dynamic-programming

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func minCostClimbingStairs(cost []int) int {
    top := len(cost)
    dp := make([]int, top+1)

    dp[0] = 0
    dp[1] = 0

    for i := 2; i < len(dp); i++ {
        dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
    }

    return dp[top]
}
