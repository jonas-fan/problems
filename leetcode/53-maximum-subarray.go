// #array

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxSubArray(nums []int) int {
    out := nums[0]
    current := nums[0]

    for i := 1; i < len(nums); i++ {
        current = nums[i] + max(current, 0)
        out = max(out, current)
    }

    return out
}
