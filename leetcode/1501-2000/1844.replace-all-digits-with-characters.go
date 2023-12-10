// #string

func replaceDigits(s string) string {
    out := []rune(s)

    for i := 0; i < len(out) - 1; i += 2 {
        out[i+1] = out[i] + out[i+1] - '0'
    }

    return string(out)
}
