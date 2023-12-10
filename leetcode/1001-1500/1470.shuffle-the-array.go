func shuffle(nums []int, n int) []int {
    out := make([]int, 0, len(nums))

    for index := 0; index < n; index++ {
        out = append(out, nums[index])
        out = append(out, nums[index+n])
    }

    return out
}
