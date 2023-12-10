// #array #dynamic-programming

func fib(n int) int {
    fib := make([]int, 0, n+1)

    fib = append(fib, 0)
    fib = append(fib, 1)

    for i := 2; i <= n; i++ {
        fib = append(fib, fib[i-1] + fib[i-2])
    }

    return fib[n]
}
