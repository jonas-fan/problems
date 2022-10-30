// #string

func checkIfPangram(sentence string) bool {
    seen := map[rune]bool{}

    for _, letter := range sentence {
        seen[letter] = true
    }

    return len(seen) == 26
}
