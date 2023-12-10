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

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func dfs(node *TreeNode) (int, int) {
    if node == nil {
        return math.MinInt, 0
    }

    lmax, lsum := dfs(node.Left)
    rmax, rsum := dfs(node.Right)

    lsum = max(lsum, 0)
    rsum = max(rsum, 0)

    out := max(max(lmax, rmax), lsum+rsum+node.Val)
    sum := max(lsum, rsum) + node.Val

    return out, sum
}

func maxPathSum(root *TreeNode) int {
    out, _ := dfs(root)

    return out
}
