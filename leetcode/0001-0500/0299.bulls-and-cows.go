// #hashmap

import (
    "fmt"
)

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    want := [10]int{}
    seen := [10]int{}

    for i := range secret {
        if secret[i] == guess[i] {
            bulls++
        } else {
            want[secret[i]-'0']++
            seen[guess[i]-'0']++
        }
    }

    for i := range seen {
        cows += min(seen[i], want[i])
    }

    return fmt.Sprintf("%dA%dB", bulls, cows)
}
