/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(lhs *TreeNode, rhs *TreeNode) bool {
    if lhs == nil && rhs == nil {
        return true
    } else if lhs == nil || rhs == nil {
        return false
    }

    if lhs.Val != rhs.Val {
        return false
    }

    return isSameTree(lhs.Left, rhs.Left) &&
           isSameTree(lhs.Right, rhs.Right)
}
