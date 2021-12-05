func rob(nums []int) int {
    length := len(nums)

    if length == 1 {
        return nums[0]
    }

    return max(max(sum(nums, 0, length-1)), max(sum(nums, 1, length)))
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func sum(nums []int, index int, length int) (int, int) {
    if index >= length {
        return 0, 0
    }

    robbed, safe := sum(nums, index+1, length)

    return (nums[index] + safe), max(robbed, safe)
}
