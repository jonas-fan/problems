// #array

func moveZeroes(nums []int)  {
    end := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            continue
        }

        nums[end] = nums[i]
        end++
    }

    for ; end < len(nums); end++ {
        nums[end] = 0
    }
}
