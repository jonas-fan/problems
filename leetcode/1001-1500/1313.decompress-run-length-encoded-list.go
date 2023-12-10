func decompressRLElist(nums []int) []int {
    size := 0

    for i := 0; i < len(nums); i += 2 {
        size += nums[i]
    }

    out := make([]int, 0, size)

    for i := 0; i < len(nums); i += 2 {
        for count := nums[i]; count > 0; count-- {
            out = append(out, nums[i+1])
        }
    }

    return out
}
