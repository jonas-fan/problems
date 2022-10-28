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
    out := [][]string{}
    signatures := map[string]int{}

    for _, str := range strs {
        signature := sortString(str)

        if pos, ok := signatures[signature]; ok {
            out[pos] = append(out[pos], str)
        } else {
            signatures[signature] = len(out)
            out = append(out, []string{str})
        }
    }

    return out
}
