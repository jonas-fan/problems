// #array

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

func countGoodRectangles(rectangles [][]int) int {
    longest := 0
    sides := map[int]int{}

    for _, rect := range rectangles {
        side := min(rect[0], rect[1])

        longest = max(longest, side)
        sides[side]++
    }

    return sides[longest]
}
