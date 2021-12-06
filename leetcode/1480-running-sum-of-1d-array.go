func runningSum(nums []int) []int {
    out := make([]int, len(nums))
    sum := 0

    for index, each := range nums {
        out[index] = each + sum
        sum = out[index]
    }

    return out
}
