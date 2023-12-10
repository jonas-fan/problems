// #string

import (
    "strings"
)

func stringMatching(words []string) []string {
    seen := map[string]bool{}

    for i := 0; i < len(words); i++ {
        if _, ok := seen[words[i]]; ok {
            continue
        }

        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }

            if strings.Contains(words[j], words[i]) {
                seen[words[i]] = true
                break;
            }
        }
    }

    out := make([]string, 0, len(seen))

    for word := range seen {
        out = append(out, word)
    }

    return out
}
