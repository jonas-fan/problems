// #string

import (
    "strings"
)

func sortSentence(s string) string {
    tokens := strings.Split(s, " ")
    out := make([]string, len(tokens))

    for _, token := range tokens {
        bound := len(token) - 1
        index := token[bound] - '1'

        out[index] = token[:bound]
    }

    return strings.Join(out, " ")
}
