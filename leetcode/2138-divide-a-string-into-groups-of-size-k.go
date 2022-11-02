// #string

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func divideString(s string, k int, fill byte) []string {
    out := make([]string, 0, len(s) / k)

    for i := 0; i < len(s); i += k {
        end := min(i+k, len(s))
        token := []rune(s[i:end])

        for len(token) < k {
            token = append(token, rune(fill))
        }

        out = append(out, string(token))
    }

    return out
}
