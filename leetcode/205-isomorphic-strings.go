// #string

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    lseen := map[byte]byte{}
    rseen := map[byte]byte{}

    for i := 0; i < len(s); i++ {
        left, right := s[i], t[i]

        if val, have := lseen[left]; have && val != right {
            return false
        }

        if val, have := rseen[right]; have && val != left {
            return false
        }

        lseen[left] = right
        rseen[right] = left
    }

    return true
}
