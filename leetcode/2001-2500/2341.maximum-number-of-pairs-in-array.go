// #array

func numberOfPairs(nums []int) []int {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num]++
    }

    removes := 0
    remains := 0

    for _, count := range seen {
        removes += count / 2
        remains += count % 2
    }

    return []int{removes, remains}
}
