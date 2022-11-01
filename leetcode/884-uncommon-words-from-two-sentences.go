// #string

import (
    "strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
    out := []string{}
    seen := map[string]int{}

    for _, str := range [2]string{s1, s2} {
        for _, token := range strings.Split(str, " ") {
            seen[token]++
        }
    }

    for token, count := range seen {
        if count == 1 {
            out = append(out, token)
        }
    }

    return out
}
