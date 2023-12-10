func decode(encoded []int, first int) []int {
    raw := make([]int, len(encoded)+1)
    raw[0] = first

    for index, each := range encoded {
        raw[index+1] = raw[index] ^ each
    }

    return raw
}
