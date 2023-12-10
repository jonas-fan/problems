// #hashmap

func checkIfExist(arr []int) bool {
    seen := map[int]int{}

    for _, num := range arr {
        if _, have := seen[num<<1]; have {
            return true
        } else if num & 0x1 == 0 {
            if _, have = seen[num>>1]; have {
                return true
            }
        }

        seen[num]++
    }

    return false
}
