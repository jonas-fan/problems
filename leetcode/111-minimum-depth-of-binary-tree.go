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

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func walk(node *TreeNode, depth int) int {
    if node == nil {
        return math.MaxInt
    }

    if node.Left == nil && node.Right == nil {
        return depth
    }

    ldepth := walk(node.Left, depth+1)
    rdepth := walk(node.Right, depth+1)

    return min(ldepth, rdepth)
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return walk(root, 1)
}
