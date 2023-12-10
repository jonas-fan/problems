// #backtracking #depth-first-search

func dup(nums []int) []int {
    return append([]int{}, nums...)
}

func walk(begin int, end int, want int, set []int, out [][]int) [][]int {
    if want == len(set) {
        return append(out, dup(set))
    } else if begin > end {
        return out
    }

    for i := begin; i <= end; i++ {
        set = append(set, i)
        out = walk(i+1, end, want, set, out)
        set = set[:len(set)-1]
    }

    return out
}

func combine(n int, k int) [][]int {
    return walk(1, n, k, []int{}, [][]int{})
}
