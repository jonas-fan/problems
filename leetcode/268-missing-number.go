// #array

func missingNumber(nums []int) int {
    target := (1 + len(nums)) * len(nums) / 2

    for _, num := range nums {
        target -= num
    }

    return target
}
