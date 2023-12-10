func maximumWealth(accounts [][]int) int {
    wealth := -1

    for _, each := range accounts {
        wealth = max(wealth, sum(each))
    }

    return wealth
}

func sum(nums []int) int {
    out := 0

    for _, each := range nums {
        out += each
    }

    return out
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}
