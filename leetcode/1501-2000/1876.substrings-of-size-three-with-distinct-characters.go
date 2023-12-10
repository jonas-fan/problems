// #string

func countGoodSubstrings(s string) int {
    out := 0

    for i := 0; i < len(s) - 2; i++ {
        if s[i] == s[i+1] {
            continue
        } else if s[i] == s[i+2] {
            continue
        } else if s[i+1] == s[i+2] {
            continue
        }

        out++
    }

    return out
}
