// #depth-first-search

func isDisinct3(one int, two int, three int) bool {
    if one == two {
        return false
    } else if one == three {
        return false
    } else if two == three {
        return false
    }

    return true
}

func dfs(nums []int, index int, triplet []int) int {
    if len(triplet) == 3 {
        one := nums[triplet[0]]
        two := nums[triplet[1]]
        three := nums[triplet[2]]

        if isDisinct3(one, two, three) {
            return 1
        }

        return 0
    }

    if index >= len(nums) {
        return 0
    }

    return dfs(nums, index+1, append(triplet, index)) +
           dfs(nums, index+1, triplet)
}

func unequalTriplets(nums []int) int {
    return dfs(nums, 0, []int{})
}
