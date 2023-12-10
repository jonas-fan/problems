// #string

func countCharacters(words []string, chars string) int {
    bucket := [26]int{}

    for _, char := range chars {
        bucket[int(char-'a')]++
    }

    out := 0

    for _, word := range words {
        needs := [26]int{}

        for _, char := range word {
            needs[int(char-'a')]++
        }

        canForm := true

        for i, count := range needs {
            if count == 0 {
                continue
            } else if count <= bucket[i] {
                continue
            }

            canForm = false
            break
        }

        if canForm {
            out += len(word)
        }
    }

    return out
}
