// #array

func min(nums []int) int {
    out := 0

    for i, num := range nums {
        if i == 0 {
            out = num
        } else if num < out {
            out = num
        }
    }

    return out
}

func max(nums []int) int {
    out := 0

    for _, num := range nums {
        if num > out {
            out = num
        }
    }

    return out
}

func findGCD(nums []int) int {
    largest := max(nums)
    smallest := min(nums)

    for {
        rest := largest % smallest

        if rest == 0 {
            break
        }

        largest = smallest
        smallest = rest
    }

    return smallest
}
