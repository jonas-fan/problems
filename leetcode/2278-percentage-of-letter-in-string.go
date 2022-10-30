// #string

func percentageLetter(s string, letter byte) int {
    count := 0

    for _, each := range s {
        if byte(each) == letter {
            count++
        }
    }

    return count * 100 / len(s)
}
