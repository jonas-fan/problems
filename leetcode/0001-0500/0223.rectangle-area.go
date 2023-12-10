// #math

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func overlapArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    if ax2 <= bx1 || bx2 <= ax1 {
        return 0
    } else if ay2 <= by1 || by2 <= ay1 {
        return 0
    }

    x1 := max(ax1, bx1)
    y1 := max(ay1, by1)
    x2 := min(ax2, bx2)
    y2 := min(ay2, by2)

    return (x2-x1) * (y2-y1)
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    aarea := (ax2-ax1) * (ay2-ay1)
    barea := (bx2-bx1) * (by2-by1)
    overs := overlapArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)

    return aarea + barea - overs
}
