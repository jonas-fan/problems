// #dynamic-programming

func tribonacci(n int) int {
    t := make([]int, 0, n+1)

    t = append(t, 0)
    t = append(t, 1)
    t = append(t, 1)

    for i := 3; i <= n; i++ {
        t = append(t, t[i-1] + t[i-2] + t[i-3])
    }

    return t[n]
}
