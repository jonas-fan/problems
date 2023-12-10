// #depth-first-search

func dfs(n int, max int, out []int) []int {
    if n > max {
        return out
    }

    out = append(out, n)

    for i := 0; i < 10; i++ {
        out = dfs(n * 10 + i, max, out)
    }

    return out
}

func lexicalOrder(n int) []int {
    out := []int{}

    for i := 0; i < 9; i++ {
        out = dfs(i+1, n, out)
    }

    return out
}
