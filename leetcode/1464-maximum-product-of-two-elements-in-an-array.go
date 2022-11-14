// #array

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxProduct(nums []int) int {
    x := max(nums[0], nums[1])
    y := min(nums[0], nums[1])

    for i := 2; i < len(nums); i++ {
        if nums[i] >= x {
            y = x
            x = nums[i]
        } else if nums[i] > y {
            y = nums[i]
        }
    }

    return (x - 1) * (y - 1)
}
