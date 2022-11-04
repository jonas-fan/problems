// #string

func counts(letters string) [26]int {
    out := [26]int{}

    for _, letter := range letters {
        out[int(letter-'a')]++
    }

    return out
}

func maxNumberOfBalloons(text string) int {
    mask := counts("balloon")
    seen := counts(text)
    out := 0

    for {
        for i := 0; i < len(seen); i++ {
            if seen[i] < mask[i] {
                return out
            }

            seen[i] -= mask[i]
        }

        out++
    }

    return out
}
