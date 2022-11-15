// #array

func smallestEqual(nums []int) int {
    index := 0
    moded := 0

    for _, num := range nums {
        if moded == num {
            return index
        }

        index++
        moded++

        if moded == 10 {
            moded = 0
        }
    }

    return -1
}
