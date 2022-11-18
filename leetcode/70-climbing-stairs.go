// #dynamic-programming

func climbStairs(n int) int {
    steps := make([]int, 0, n+1)

    steps = append(steps, 1)
    steps = append(steps, 1)

    for i := 2; i <= n; i++ {
        steps = append(steps, steps[i-1] + steps[i-2])
    }

    return steps[n]
}
