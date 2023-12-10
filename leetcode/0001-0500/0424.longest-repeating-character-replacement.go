// #string

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func popular(nums []int) int {
    out := 0

    for _, num := range nums {
        out = max(out, num)
    }

    return out
}

func characterReplacement(s string, k int) int {
    out := 0
    seen := [26]int{}
    begin := 0
    end := 0

    for ; end < len(s); end++ {
        seen[s[end]-'A']++

        length := end - begin + 1

        if length - popular(seen[:]) <= k {
            out = max(out, length)
        } else {
            seen[s[begin]-'A']--
            begin++
        }
    }

    return out
}
