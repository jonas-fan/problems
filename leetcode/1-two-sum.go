func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)

    for lhs, num := range nums {
        if rhs, ok := indices[target-num]; ok {
            return []int{lhs, rhs}
        }

        indices[num] = lhs
    }

    return []int{}
}
