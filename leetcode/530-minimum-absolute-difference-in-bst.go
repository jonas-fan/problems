/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "math"
)

func abs(val int) int {
    if val < 0 {
        return -val
    }

    return val
}

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func dfs(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    out = dfs(node.Left, out)
    out = append(out, node.Val)
    out = dfs(node.Right, out)

    return out
}

func getMinimumDifference(root *TreeNode) int {
    diff := math.MaxInt
    vals := dfs(root, []int{})
    
    for i := 0; i < len(vals) - 1; i++ {
        diff = min(diff, abs(vals[i]-vals[i+1]))
    }

    return diff
}
