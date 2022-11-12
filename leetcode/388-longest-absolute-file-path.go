// #string #stack

import (
    "strings"
)

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func containsFile(dentries []string) bool {
    if len(dentries) == 0 {
        return false
    }

    last := dentries[len(dentries)-1]

    return strings.Contains(last, ".")
}

func lengthLongestPath(input string) int {
    out := 0
    stack := []string{}

    for _, token := range strings.Split(input, "\n") {
        depth := strings.Count(token, "\t")
        entry := token[depth:]

        if depth < len(stack) {
            if containsFile(stack) {
                fullpath := strings.Join(stack, "/")
                out = max(out, len(fullpath))
            }

            for depth < len(stack) {
                stack = stack[:len(stack)-1]
            }
        }

        stack = append(stack, entry)
    }

    if containsFile(stack) {
        fullpath := strings.Join(stack, "/")
        out = max(out, len(fullpath))
    }

    return out
}
