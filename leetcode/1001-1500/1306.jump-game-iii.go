func canReach(arr []int, start int) bool {
    visited := make(map[int]bool, len(arr))

    return reachable(arr, &visited, start)
}

func reachable(arr []int, visited *map[int]bool, start int) bool {
    if start < 0 || start >= len(arr) {
        return false
    } else if (*visited)[start] {
        return false
    }

    val := arr[start]

    if val == 0 {
        return true
    }

    (*visited)[start] = true

    return reachable(arr, visited, start-val) || reachable(arr, visited, start+val)
}
