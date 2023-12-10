// #array

func pivotIndex(nums []int) int {
    left := 0
    right := 0

    for _, num := range nums {
        right += num
    }

    for i := 0; i < len(nums); i++ {
        right -= nums[i]

        if left == right {
            return i
        }

        left += nums[i]
    }

    return -1
}
