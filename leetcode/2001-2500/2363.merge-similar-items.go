// #array

import (
    "sort"
)

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    values := make([]int, 0, max(len(items1), len(items2)))
    bucket := map[int]int{}

    for _, item := range items1 {
        if _, have := bucket[item[0]]; !have {
            values = append(values, item[0])
        }

        bucket[item[0]] += item[1]
    }

    for _, item := range items2 {
        if _, have := bucket[item[0]]; !have {
            values = append(values, item[0])
        }

        bucket[item[0]] += item[1]
    }

    sort.Ints(values)

    out := make([][]int, 0, len(values))

    for _, value := range values {
        out = append(out, []int{value, bucket[value]})
    }

    return out
}
