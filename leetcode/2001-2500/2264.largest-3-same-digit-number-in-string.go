// #string

func allEqual(letters []rune) bool {
    for i := 0; i < len(letters) - 1; i++ {
        if letters[i] != letters[i+1] {
            return false
        }
    }

    return true
}

func largestGoodInteger(num string) string {
    out := ""
    digits := []rune(num)

    for i := 0; i < len(digits) - 2; i++ {
        if allEqual(digits[i:i+3]) {
            str := string(digits[i:i+3])

            if str > out {
                out = str
            }
        }
    }

    return out
}
