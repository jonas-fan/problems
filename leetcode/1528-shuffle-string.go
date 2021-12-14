func restoreString(s string, indices []int) string {
    plaintext := make([]byte, len(indices))

    for index, each := range indices {
        plaintext[each] = s[index]
    }

    return string(plaintext)
}
