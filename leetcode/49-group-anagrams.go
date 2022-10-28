import (
    "sort"
)

func sortString(str string) string {
    out := []rune(str)

    sort.Slice(out, func(lhs, rhs int) bool {
        return out[lhs] < out[rhs]
    })

    return string(out)
}

func groupAnagrams(strs []string) [][]string {
    buckets := map[string][]string{}

    for _, str := range strs {
        sorted := sortString(str)

        buckets[sorted] = append(buckets[sorted], str)
    }

    out := make([][]string, 0, len(buckets))

    for _, bucket := range buckets {
        out = append(out, bucket)
    }

    return out
}
