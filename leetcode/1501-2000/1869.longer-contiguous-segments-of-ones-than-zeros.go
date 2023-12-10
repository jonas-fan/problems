// #string

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxContiguous(chars []rune,  target rune) int {
    out := 0
    count := 0

    for i := 0; i < len(chars) - 1; i++ {
        if chars[i] == target && chars[i] == chars[i+1] {
            count++
        } else {
            out = max(out, count + 1)
            count = 0
        }
    }

    if chars[len(chars)-1] == target {
        count++
    }

    return max(out, count)
}

func checkZeroOnes(s string) bool {
    chars := []rune(s)

    return maxContiguous(chars, '1') > maxContiguous(chars, '0')
}
