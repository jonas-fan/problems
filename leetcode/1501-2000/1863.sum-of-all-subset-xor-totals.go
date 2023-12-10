// #array

func walk(nums []int, xorSum int) int {
    if len(nums) == 0 {
        return xorSum
    }

    return walk(nums[1:], xorSum ^ nums[0]) + walk(nums[1:], xorSum)
}

func subsetXORSum(nums []int) int {
    return walk(nums, 0)
}
