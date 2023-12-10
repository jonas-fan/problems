// #array

func sortArrayByParity(nums []int) []int {
    end := 0

    for i, num := range nums {
        if num & 0x1 == 0 {
            nums[end], nums[i] = nums[i], nums[end]
            end++
        }
    }

    return nums
}
