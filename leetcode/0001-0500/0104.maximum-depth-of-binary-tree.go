// #tree #binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
