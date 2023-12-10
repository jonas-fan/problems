// #bits

func singleNumber(nums []int) int {
    out := 0

    for _, num := range nums {
        out = out ^ num
    }

    return out
}
