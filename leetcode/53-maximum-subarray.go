// #array

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxSubArray(nums []int) int {
    sum := make([]int, len(nums))
    sum[0] = nums[0]

    out := sum[0]

    for i := 1; i < len(nums); i++ {
        sum[i] = nums[i] + max(sum[i-1], 0)
        out = max(out, sum[i])
    }

    return out
}
