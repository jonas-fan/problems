// #array

func runningSum(nums []int) []int {
    out := make([]int, len(nums))

    if len(out) > 0 {
        out[0] = nums[0]

        for i := 1; i < len(out); i++ {
            out[i] = nums[i] + out[i-1]
        }
    }

    return out
}
