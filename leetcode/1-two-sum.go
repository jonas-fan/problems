// #array #hashmap

func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, num := range nums {
        rest := target - num

        if where, have := seen[rest]; have {
            return []int{i, where}
        }

        seen[num] = i
    }

    return []int{}
}
