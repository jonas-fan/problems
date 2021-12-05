func lengthOfLongestSubstring(s string) int {
    longest := 0
    begin := -1
    seen := map[rune]int{}

    for index, each := range s {
        if at, ok := seen[each]; ok {
            begin = max(begin, at)
        }

        seen[each] = index
        longest = max(longest, index-begin)
    }

    return longest
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}
