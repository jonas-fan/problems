// #dynamic-programming

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    } else if len(nums) == 1 {
        return nums[0]
    }

    money := make([]int, 0, len(nums))

    money = append(money, nums[0])
    money = append(money, max(nums[0], nums[1]))

    for i := 2; i < len(nums); i++ {
        money = append(money, max(money[i-1], money[i-2]+nums[i]))
    }

    return money[len(nums)-1]
}
