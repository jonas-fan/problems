// #string

func eval(str string) int {
    out := 0

    for _, letter := range str {
        out = out * 10 + (int(letter) - 'a')
    }

    return out
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    sum := eval(firstWord) + eval(secondWord)

    return sum == eval(targetWord)
}
