// #array

func sumZero(n int) []int {
    out := make([]int, 0, n)

    if n & 0x1 == 1 {
        out = append(out, 0)
        n--
    }

    for n > 0 {
        out = append(out, n)
        out = append(out, -n)
        n -= 2
    }

    return out
}
