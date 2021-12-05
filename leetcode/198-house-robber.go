func rob(nums []int) int {
    return max(sum(nums, 0, len(nums)))
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func sum(nums []int, index int, length int) (int, int) {
    if index == length {
        return 0, 0
    }

    robbed, safe := sum(nums, index+1, length)

    return (nums[index] + safe), max(robbed, safe)
}
