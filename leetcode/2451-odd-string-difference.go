// #string

func oddString(words []string) string {
    if len(words) == 0 {
        return ""
    }

    diffs := make([]int, len(words))

    for j := 1; j < len(words[0]); j++ {
        for i := 0; i < len(words); i++ {
            diffs[i] = int(words[i][j] - words[i][j-1])
        }

        for i := 0; i < len(diffs) - 2; i++ {
            if diffs[i] == diffs[i+1] && diffs[i] == diffs[i+2] {
                continue
            }

            if diffs[i] == diffs[i+1] {
                return words[i+2]
            } else if diffs[i] == diffs[i+2] {
                return words[i+1]
            } else {
                return words[i]
            }
        }
    }

    return ""
}
