// #backtracking

import (
    "fmt"
    "sort"
)

func dup(nums []int) []int {
    return append([]int{}, nums...)
}

func dfs(candidates []int, target int, nums []int, out [][]int) [][]int {
    if target < 0 {
        return out
    } else if target == 0 {
        return append(out, dup(nums))
    }

    for i, num := range candidates {
        if i > 0 && candidates[i] == candidates[i-1] {
            continue
        }

        nums = append(nums, num)
        out = dfs(candidates[i+1:], target-num, nums, out)
        nums = nums[:len(nums)-1]
    }

    return out
}

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)

    return dfs(candidates, target, []int{}, [][]int{})
}
