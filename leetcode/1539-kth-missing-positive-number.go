// #binary-search

func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)

    for left < right {
        mid := left + (right - left) >> 1

        if (arr[mid] - mid) > k {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return k + left
}
