// #array

func repeatedNTimes(nums []int) int {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num]++

        if seen[num] > 1 {
            return num
        }
    }

    return 0
}
