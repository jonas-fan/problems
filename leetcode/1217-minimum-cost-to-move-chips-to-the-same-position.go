func minCostToMoveChips(position []int) int {
    chips := len(position)
    odd := 0

    for _, pos := range position {
        odd += pos & 0x1
    }

    if odd <= (chips >> 1) {
        return odd
    }

    return chips-odd
}
