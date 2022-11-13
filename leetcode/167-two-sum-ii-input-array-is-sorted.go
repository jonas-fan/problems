// #array

func twoSum(numbers []int, target int) []int {
    seen := map[int]int{}

    for i, num := range numbers {
        rest := target - num

        if where, have := seen[rest]; have {
            return []int{where, i+1}
        }

        seen[num] = i + 1
    }

    return []int{}
}
