// #string

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func lengthOfLongestSubstring(s string) int {
    out := 0
    begin := -1
    seen := map[rune]int{}

    for i, char := range s {
        if where, have := seen[char]; have {
            begin = max(begin, where)
        }

        seen[char] = i
        out = max(out, i-begin)
    }

    return out
}
