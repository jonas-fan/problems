// #array

func removeDuplicates(nums []int) int {
    size := 0

    for i := 0; i < len(nums); i++ {
        if size > 0 && nums[i] == nums[size-1] {
            continue
        }

        nums[size] = nums[i]
        size++
    }

    return size
}
