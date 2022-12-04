// #binary-search

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func findMin(nums []int) int {
    left := 0
    right := len(nums)

    for left < right {
        mid := left + (right - left) >> 1

        if nums[left] <= nums[mid] && nums[mid] <= nums[right-1] {
            return nums[left]
        } else if nums[mid] < nums[mid-1] {
            return nums[mid]
        }

        if nums[left] < nums[mid] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[min(left, len(nums)-1)]
}
