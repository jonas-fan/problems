// #string

func countAsterisks(s string) int {
    asterisks := 0
    counting := true

    for _, letter := range s {
        if letter == '|' {
            counting = !counting
            continue
        }

        if counting && letter == '*' {
            asterisks++
        }
    }

    return asterisks
}
