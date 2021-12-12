func canPartition(nums []int) bool {
    sum := 0

    for _, each := range nums {
        sum += each
    }

    if sum&0x1 != 0 {
        return false
    }

    return dfs(nums, 0, len(nums), sum>>1, &map[int]bool{})
}

func dfs(nums []int, pos int, length int, target int, cache *map[int]bool) bool {
    if target == 0 {
        return true
    } else if target < 0 {
        return false
    } else if pos >= length {
        return false
    }

    index := key(pos, target)

    if ok, have := (*cache)[index]; have {
        return ok
    }

    (*cache)[index] = dfs(nums, pos+1, length, target, cache) ||
        dfs(nums, pos+1, length, target-nums[pos], cache)

    return (*cache)[index]
}

func key(pos int, target int) int {
    return pos<<16 + target
}
