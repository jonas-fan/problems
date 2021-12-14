func numJewelsInStones(jewels string, stones string) int {
    box := make(map[rune]int)

    for _, each := range stones {
        box[each]++
    }

    count := 0

    for _, each := range jewels {
        count += box[each]
    }

    return count
}
