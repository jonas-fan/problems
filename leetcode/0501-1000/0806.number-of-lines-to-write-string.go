// #string

func numberOfLines(widths []int, s string) []int {
    line := 1
    pixels := 0

    for _, c := range s {
        width := widths[int(c-'a')]

        if pixels + width > 100 {
            line++
            pixels = 0
        }

        pixels += width
    }

    return []int{line, pixels}
}
