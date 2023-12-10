// #string

import (
    "strings"
    "strconv"
)

func ascending(nums []int) bool {
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] >= nums[i+1] {
            return false
        }
    }

    return true
}

func areNumbersAscending(s string) bool {
    words := strings.Split(s, " ")
    nums := []int{}

    for _, word := range words {
        if num, err := strconv.Atoi(word); err == nil {
            nums = append(nums, num) 
        }
    }

    return ascending(nums)
}
