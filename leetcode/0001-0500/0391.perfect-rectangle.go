// #array

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

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

func area(rect []int) int {
    return (abs(rect[0]-rect[2])) * (abs(rect[1]-rect[3]))
}

func isOverlapped(lhs []int, rhs[]int) bool {
    if lhs[0] >= rhs[2] || lhs[2] <= rhs[0] {
        return false
    }

    if lhs[1] >= rhs[3] || lhs[3] <= rhs[1] {
        return false
    }

    return true
}

func isRectangleCover(rectangles [][]int) bool {
    areaTotal := 0
    merged := []int{rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]}

    for i := 0; i < len(rectangles); i++ {
        areaTotal += area(rectangles[i])
        merged[0] = min(merged[0], rectangles[i][0])
        merged[1] = min(merged[1], rectangles[i][1])
        merged[2] = max(merged[2], rectangles[i][2])
        merged[3] = max(merged[3], rectangles[i][3])

        for j := i + 1; j < len(rectangles); j++ {
            if isOverlapped(rectangles[i], rectangles[j]) {
                return false
            }
        }
    }

    return area(merged) == areaTotal
}
