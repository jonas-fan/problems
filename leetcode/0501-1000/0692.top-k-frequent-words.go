// #array

import (
    "sort"
)

func topKFrequent(words []string, k int) []string {
    seen := map[string]int{}

    for _, word := range words {
        seen[word]++
    }

    words = words[:0]

    for word := range seen {
        words = append(words, word)
    }

    sort.Slice(words, func(lhs int, rhs int) bool {
        if seen[words[lhs]] == seen[words[rhs]] {
            return words[lhs] < words[rhs]
        }

        return seen[words[lhs]] > seen[words[rhs]]
    })

    return words[:k]
}
