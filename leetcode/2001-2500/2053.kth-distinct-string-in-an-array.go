// #string

func kthDistinct(arr []string, k int) string {
    seen := map[string]int{}

    for _, str := range arr {
        seen[str]++
    }

    for _, str := range arr {
        if seen[str] > 1 {
            continue
        }

        if k--; k == 0 {
            return str
        }
    }

    return ""
}
