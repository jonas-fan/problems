// #backtracking

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
        nums = append(nums, num)
        out = dfs(candidates[i:], target-num, nums, out)
        nums = nums[:len(nums)-1]
    }

    return out
}

func combinationSum(candidates []int, target int) [][]int {
    return dfs(candidates, target, []int{}, [][]int{})
}
