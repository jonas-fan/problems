// #array

func minOperations(nums []int) int {
    out := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            continue
        }

        out += nums[i-1] - nums[i] + 1
        nums[i] = nums[i-1] + 1
    }

    return out
}
