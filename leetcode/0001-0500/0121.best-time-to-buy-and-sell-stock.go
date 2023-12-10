// #array

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxProfit(prices []int) int {
    out := 0
    peak := 0

    for i := len(prices) - 1; i >= 0; i-- {
        out = max(out, peak - prices[i])
        peak = max(peak, prices[i])
    }

    return out
}
