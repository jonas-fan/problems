func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    for {
        if len(word1) == 0 && len(word2) == 0 {
            return true
        } else if len(word1) == 0 || len(word2) == 0 {
            return false
        }

        if len(word1[0]) == 0 || len(word2[0]) == 0 {
            if len(word1[0]) == 0 {
                word1 = word1[1:]
            }

            if len(word2[0]) == 0 {
                word2 = word2[1:]
            }

            continue
        }

        if word1[0][0] != word2[0][0] {
            return false
        }

        word1[0] = word1[0][1:]
        word2[0] = word2[0][1:]
    }
}
