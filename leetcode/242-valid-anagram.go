// #string

func isAnagram(s string, t string) bool {
    lhs := [26]int{}
    rhs := [26]int{}

    for _, letter := range s {
        lhs[int(letter-'a')]++
    }

    for _, letter := range t {
        rhs[int(letter-'a')]++
    }

    for i := 0; i < len(lhs); i++ {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}
