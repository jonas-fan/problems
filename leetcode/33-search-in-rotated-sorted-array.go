// #binary-search

func search(nums []int, target int) int {
    left := 0
    right := len(nums)

    for left < right {
        mid := left + (right - left) >> 1

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            if nums[left] < nums[mid] && nums[mid] < nums[right-1] {
                left = mid + 1
            } else if nums[left] < nums[mid] {
                left = mid + 1
            } else if target <= nums[right-1] {
                left = mid + 1
            } else {
                right = mid
            }
        } else {
            if nums[left] < nums[mid] && nums[mid] < nums[right-1] {
                right = mid
            } else if nums[mid] < nums[right-1] {
                right = mid
            } else if nums[left] <= target {
                right = mid
            } else {
                left = mid + 1
            }
        }
    }

    return -1
}
