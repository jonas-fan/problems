func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0

    for _, each := range candies {
        if each > max {
            max = each
        }
    }

    bound := max - extraCandies
    out := make([]bool, 0, len(candies))

    for _, each := range candies {
        out = append(out, each >= bound)
    }

    return out
}
