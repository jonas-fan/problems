// #array

func counts(nums []int) map[int]int {
    out := map[int]int{}

    for _, num := range nums {
        out[num]++
    }

    return out
}

func equals(lhs map[int]int, rhs map[int]int) bool {
    if len(lhs) != len(rhs) {
        return false
    }

    for key := range lhs {
        if lhs[key] != rhs[key] {
            return false
        }
    }

    return true
}

func canBeEqual(target []int, arr []int) bool {
    return equals(counts(target), counts(arr))
}
