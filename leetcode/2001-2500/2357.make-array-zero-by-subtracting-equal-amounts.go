// #array

func minimumOperations(nums []int) int {
    seen := map[int]int{}

    for _, num := range nums {
        if num != 0 {
            seen[num]++
        }
    }

    return len(seen)
}
