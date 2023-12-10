func getConcatenation(nums []int) []int {
    length := len(nums)
    out := make([]int, length*2)

    for index, each := range nums {
        out[index] = each
        out[index+length] = each
    }

    return out
}
