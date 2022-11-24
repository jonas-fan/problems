// #array #binary-search

func peakIndexInMountainArray(arr []int) int {
    left := 0
    right := len(arr) - 1

    for right - left > 2 {
        mid := (left + right) >> 1

        if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
            return mid
        } else if arr[mid-1] < arr[mid] {
            left = mid
        } else if arr[mid] > arr[mid+1] {
            right = mid
        }
    }

    return left + 1
}
