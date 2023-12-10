// #string

func isSorted(strs []string, col int) bool {
    for i := 0; i < len(strs) - 1; i++ {
        if strs[i][col] <= strs[i+1][col] {
            // ascending
        } else {
            return false
        }
    }

    return true
}

func minDeletionSize(strs []string) int {
    out := 0
    
    if len(strs) == 0 {
        return 0
    }

    for i := range strs[0] {
        if !isSorted(strs, i) {
            out++
        }
    }

    return out
}
