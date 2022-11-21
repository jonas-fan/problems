// #depth-first-search

func dup(nums []int) []int {
    return append([]int{}, nums...)
}

func walk(nums []int, bucket []int, out[][]int) [][]int {
    if len(nums) == 0 {
        return append(out, dup(bucket))
    }

    for i := 0; i < len(nums); i++ {
        nextNums := append(dup(nums[:i]), nums[i+1:]...)
        nextBucket := append(bucket, nums[i])

        out = walk(nextNums, nextBucket, out)
    }

    return out
}

func permute(nums []int) [][]int {
    return walk(nums, []int{}, [][]int{})
}
