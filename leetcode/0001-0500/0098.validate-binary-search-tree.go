// #tree #binary-tree #binary-search-tree

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

func walk(node *TreeNode, min int, max int) bool {
    if node == nil {
        return true
    }

    if node.Val <= min {
        return false
    } else if node.Val >= max {
        return false
    }

    return walk(node.Left, min, node.Val) &&
           walk(node.Right, node.Val, max)
}

func isValidBST(root *TreeNode) bool {
    return walk(root, math.MinInt, math.MaxInt)
}
