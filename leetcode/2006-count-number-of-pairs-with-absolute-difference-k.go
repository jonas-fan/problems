// #array

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func countKDifference(nums []int, k int) int {
    out := 0

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if abs(nums[i] - nums[j]) == k {
                out++
            }
        }
    }

    return out
}
