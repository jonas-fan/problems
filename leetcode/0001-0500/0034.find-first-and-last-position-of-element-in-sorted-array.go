// #binary-search

func findFirst(nums []int, target int) int {
    left := 0
    right := len(nums)

    for left < right {
        mid := left + (right - left) >> 1

        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }

    if left == len(nums) || nums[left] != target {
        return -1
    }

    return left
}

func findLast(nums []int, target int) int {
    left := 0
    right := len(nums)

    for left < right {
        mid := left + (right - left) >> 1

        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    if left == 0 || nums[left-1] != target {
        return -1
    }

    return left - 1
}

func searchRange(nums []int, target int) []int {
    return []int{
        findFirst(nums, target),
        findLast(nums, target),
    }
}
