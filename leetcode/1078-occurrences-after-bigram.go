// #string

import (
    "strings"
)

func findOcurrences(text string, first string, second string) []string {
    out := []string{}
    words := strings.Split(text, " ")

    for i := 0; i < len(words) - 2; i++ {
        if words[i] != first {
            continue
        } else if words[i+1] != second {
            continue
        }

        out = append(out, words[i+2])
    }

    return out
}
