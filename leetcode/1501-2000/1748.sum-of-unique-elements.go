// #array

func sumOfUnique(nums []int) int {
    sum := 0
    seen := map[int]int{}

    for _, num := range nums {
        if count, have := seen[num]; !have {
            sum += num
        } else if count == 1 {
            sum -= num
        }

        seen[num]++
    }

    return sum
}
