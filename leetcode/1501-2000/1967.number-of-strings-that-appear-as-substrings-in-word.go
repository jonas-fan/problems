// #string

import (
    "strings"
)

func numOfStrings(patterns []string, word string) int {
    found := 0

    for _, pattern := range patterns {
        if strings.Contains(word, pattern) {
            found++
        }
    }

    return found
}
