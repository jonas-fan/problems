// #array

import (
    "sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
    index := 0

    for _, target := range arr2 {
        for i := index; i < len(arr1); i++ {
            if arr1[i] == target {
                arr1[i], arr1[index] = arr1[index], target
                index++
            }
        }
    }

    sort.Ints(arr1[index:])

    return arr1
}
