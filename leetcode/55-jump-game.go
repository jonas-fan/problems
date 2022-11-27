// #dynamic-programming

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func canJump(nums []int) bool {
    longest := 0

    for i, num := range nums {
        if longest < i {
            return false
        } else if longest + 1 >= len(nums) {
            return true
        }

        longest = max(longest, i+num)
    }

    return false
}
