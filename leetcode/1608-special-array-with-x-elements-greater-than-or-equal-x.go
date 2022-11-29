import (
    "sort"
)

func specialArray(nums []int) int {
    sort.Ints(nums)

    left := 0
    right := len(nums)

    for left < right {
        mid := left + (right - left) >> 1
        cnt := len(nums) - mid

        if nums[mid] >= cnt {
            if mid == 0 || nums[mid-1] < cnt {
                return cnt
            }

            right = mid
        } else {
            left = mid + 1
        }
    }

    return -1
}
