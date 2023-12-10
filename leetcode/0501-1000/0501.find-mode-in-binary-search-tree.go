/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "sort"
)

func walk(node *TreeNode, nums map[int]int) map[int]int {
    if node == nil {
        return nums
    }

    nums[node.Val]++

    nums = walk(node.Left, nums)
    nums = walk(node.Right, nums)

    return nums
}

func findMode(root *TreeNode) []int {
    seen := walk(root, map[int]int{})
    keys := make([]int, 0, len(seen))

    for key := range seen {
        keys = append(keys, key)
    }

    sort.Slice(keys, func(lhs int, rhs int) bool {
        return seen[keys[lhs]] > seen[keys[rhs]]
    })

    for i := 1; i < len(keys); i++ {
        if seen[keys[i]] < seen[keys[i-1]] {
            return keys[0:i]            
        }
    }

    return keys
}
