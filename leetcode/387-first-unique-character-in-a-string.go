// #string

func firstUniqChar(s string) int {
    seen := map[rune]int{}

    for _, char := range s {
        seen[char]++
    }

    for i, char := range s {
        if seen[char] == 1 {
            return i
        }
    }

    return -1;
}
