func numIdenticalPairs(nums []int) int {
    out := 0

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                out++
            }
        }
    }

    return out
}
