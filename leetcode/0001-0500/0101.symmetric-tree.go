// #tree #binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(lhs *TreeNode, rhs *TreeNode) bool {
    if lhs == nil && rhs == nil {
        return true
    } else if lhs == nil {
        return false
    } else if rhs == nil {
        return false
    } else if lhs.Val != rhs.Val {
        return false
    }

    return walk(lhs.Left, rhs.Right) &&
           walk(lhs.Right, rhs.Left)
}

func isSymmetric(root *TreeNode) bool {
    return walk(root.Left, root.Right)
}
