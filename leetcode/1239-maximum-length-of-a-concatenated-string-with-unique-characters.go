func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func hash(token string) int {
    out := 0

    for _, letter := range token {
        bit := 0x1 << int(letter - 'a')

        // conflict?
        if out & bit > 0 {
            return 0
        }

        out = out | bit
    }

    return out
}

func dfs(tokens []string, signatures []int, length int, signature int) int {
    out := length

    for i := range tokens {
        if signatures[i] == 0 {
            continue
        }

        result := 0

        if signature & signatures[i] > 0 {
            result = dfs(tokens[i+1:], signatures[i+1:], len(tokens[i]), signatures[i])
        } else {
            result = dfs(tokens[i+1:], signatures[i+1:], len(tokens[i]) + length, signatures[i] | signature)
        }

        out = max(out, result)
    }

    return out
}

func maxLength(arr []string) int {
    signatures := make([]int, len(arr))

    for i, token := range arr {
        signatures[i] = hash(token)
    }

    return dfs(arr, signatures, 0, 0)
}
