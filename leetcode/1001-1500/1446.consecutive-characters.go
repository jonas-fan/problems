func maxPower(s string) int {
    var out int
    var count int
    var last rune

    for _, each := range s {
        if each != last {
            out = max(out, count)
            last = each
            count = 1
        } else {
            count++
        }
    }

    return max(out, count)
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}
