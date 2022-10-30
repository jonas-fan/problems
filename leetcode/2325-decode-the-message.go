// #string

func decodeMessage(key string, message string) string {
    table := map[rune]rune{}

    for _, letter := range key {
        if letter == ' ' {
            continue
        } else if _, ok := table[letter]; ok {
            continue
        }

        table[letter] = 'a' + rune(len(table))
    }

    out := make([]rune, 0, len(message))

    for _, letter := range message {
        if letter == ' ' {
            out = append(out, letter)
        } else {
            out = append(out, table[letter])
        }
    }

    return string(out)
}
