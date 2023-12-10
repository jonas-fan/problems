// #dynamic-programming

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func jump(nums []int) int {
    out := 0
    longest := 0
    offset := 0

    for i := 0; i < len(nums) - 1; i++ {
        longest = max(longest, i+nums[i])

        if i == offset {
            offset = longest
            out++
        }
    }

    return out
}
