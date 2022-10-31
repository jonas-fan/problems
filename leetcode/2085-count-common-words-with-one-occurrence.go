// #string

func counts(words []string) map[string]int {
    out := map[string]int{}

    for _, word := range words {
        out[word]++
    }

    return out
}

func countWords(words1 []string, words2 []string) int {
    out := 0
    seen1 := counts(words1)
    seen2 := counts(words2)

    for word, count1 := range seen1 {
        if count1 > 1 {
            continue
        } else if count2, ok := seen2[word]; !ok {
            continue
        } else if count2 != 1 {
            continue
        }

        out++
    }

    return out
}
