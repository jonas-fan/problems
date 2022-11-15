// #array

func divideArray(nums []int) bool {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num]++
    }

    for _, count := range seen {
        if count & 0x1 == 1 {
            return false
        }
    }

    return true
}
