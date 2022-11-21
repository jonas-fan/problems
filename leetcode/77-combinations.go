// #depth-first-search

func dup(nums []int) []int {
    return append([]int{}, nums...)
}

func walk(num int, max int, nums int, set []int, out [][]int) [][]int {
    if nums == 0 {
        return append(out, dup(set))
    } else if num > max {
        return out
    }

    out = walk(num+1, max, nums-1, append(set, num), out)
    out = walk(num+1, max, nums, set, out)

    return out
}

func combine(n int, k int) [][]int {
    return walk(1, n, k, []int{}, [][]int{})
}
