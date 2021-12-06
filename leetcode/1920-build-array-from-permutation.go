func buildArray(nums []int) []int {
    out := make([]int, len(nums))

    for index, each := range nums {
        out[index] = nums[each]
    }

    return out
}
