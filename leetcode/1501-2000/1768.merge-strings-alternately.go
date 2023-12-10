// #string

func mergeAlternately(word1 string, word2 string) string {
    out := make([]rune, 0, len(word1)+len(word2))
    words := []string{word1, word2}
    index := 0

    for len(words[0]) != 0 || len(words[1]) != 0 {
        if len(words[index]) == 0 {
            index = 1 - index
        }

        out = append(out, rune(words[index][0]))
        words[index] = words[index][1:]
        index = 1 - index
    }

    return string(out)
}
