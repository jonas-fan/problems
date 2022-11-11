// #search #binary-search

func search(nums []int, target int) int {
    left := 0
    right := len(nums)

    for left < right {
        mid := (left + right) >> 1

        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid
        } else {
            return mid
        }
    }

    return -1
}
