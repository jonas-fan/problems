func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func lengthOfLongestSubstring(s string) int {
    out := 0
    begin := -1
    seen := make(map[rune]int)

    for index, letter := range s {
        if where, ok := seen[letter]; ok {
            begin = max(begin, where)
        }

        seen[letter] = index
        out = max(out, index-begin)
    }

    return out
}
