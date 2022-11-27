// #binary-search

import (
    "sort"
)

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    out := 0

    sort.Ints(arr2)

    for i := 0; i < len(arr1); i++ {
        left := 0
        right := len(arr2)

        for left < right {
            mid := left + (right - left) >> 1
            diff := arr1[i] - arr2[mid]

            if abs(diff) <= d {
                break
            } else if diff < 0 {
                right = mid
            } else {
                left = mid + 1
            }
        }

        if left == right {
            out++
        }
    }

    return out
}
