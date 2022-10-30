// #string

func repeatedCharacter(s string) byte {
    bucket := make([]int, 26)

    for _, letter := range s {
        index := letter - 'a'

        if bucket[index] == 1 {
            return byte(letter)
        }

        bucket[index]++
    }

    return byte('?')
}
