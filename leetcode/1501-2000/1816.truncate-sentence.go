// #string

func truncateSentence(s string, k int) string {
    for i, letter := range s {
        if letter == ' ' {
            if k--; k == 0 {
                return s[:i]                
            }
        }
    }

    return s
}
