// #string

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func minOperations(logs []string) int {
    depth := 0

    for _, folder := range logs {
        switch folder {
        case "./":
        case "../":
            depth = max(depth - 1, 0)
        default:
            depth++
        }
    }

    return depth
}
