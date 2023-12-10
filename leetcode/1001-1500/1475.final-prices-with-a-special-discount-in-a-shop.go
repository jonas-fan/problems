// #array

func finalPrices(prices []int) []int {
    out := make([]int, len(prices))

    for i := 0; i < len(prices); i++ {
        j := i + 1

        for ; j < len(prices); j++ {
            if prices[j] <= prices[i] {
                break
            }
        }

        if j == len(prices) {
            out[i] = prices[i]
        } else {
            out[i] = prices[i] - prices[j]
        }
    }

    return out
}
