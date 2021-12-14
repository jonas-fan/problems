func smallerNumbersThanCurrent(nums []int) []int {
    bucket := make([]int, 101)

    for _, each := range nums {
        bucket[each]++
    }

    sum := 0

    for index, each := range bucket {
        bucket[index] = sum
        sum += each
    }

    out := make([]int, len(nums))

    for index, each := range nums {
        out[index] = bucket[each]
    }

    return out
}
