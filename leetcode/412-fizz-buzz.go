// #string

func fizzBuzz(n int) []string {
    out := make([]string, 0, n)

    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            out = append(out, "FizzBuzz")
        } else if i % 5 == 0 {
            out = append(out, "Buzz")
        } else if i % 3 == 0 {
            out = append(out, "Fizz")
        } else {
            out = append(out, strconv.Itoa(i))
        }
    }

    return out
}
