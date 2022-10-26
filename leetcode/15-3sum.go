import (
    "sort"
)

func threeSum(nums []int) [][]int {
    out := [][]int{}
    seen := map[[2]int]bool{}

    sort.Ints(nums)

    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            key := [2]int{nums[i], nums[j]}

            if _, ok := seen[key]; ok {
                continue
            }

            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    triplet := []int{nums[i], nums[j], nums[k]}

                    keys := [3][2]int{
                        [2]int{triplet[0], triplet[1]},
                        [2]int{triplet[0], triplet[2]},
                        [2]int{triplet[1], triplet[2]},
                    }

                    if _, ok := seen[keys[0]]; ok {
                        continue
                    } else if _, ok := seen[keys[1]]; ok {
                        continue
                    } else if _, ok := seen[keys[2]]; ok {
                        continue
                    }

                    seen[keys[0]] = true
                    seen[keys[1]] = true
                    seen[keys[2]] = true
                    out = append(out, triplet)
                }
            }
        }
    }

    return out
}
