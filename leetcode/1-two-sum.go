func twoSum(nums []int, target int) []int {
    mapping := map[int]int{}

    for index, each := range nums {
        if val, ok := mapping[target-each]; ok {
            return []int{index, val}
        }

        mapping[each] = index
    }

    return []int{0, 0}
}
