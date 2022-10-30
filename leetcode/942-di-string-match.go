// #string

func diStringMatch(s string) []int {
    out := make([]int, 0, len(s)+1)
    max := len(s)
    min := 0

    for i := 0; i < len(s); i++ {
        if s[i] == 'I' {
            out = append(out, min)
            min++
        } else {
            out = append(out, max)
            max--
        }
    }

    out = append(out, min)

    return out
}
