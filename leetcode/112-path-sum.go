/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode, target int) bool {
    if node == nil {
        return false
    }

    if node.Left == nil && node.Right == nil {
        return node.Val == target
    }

    if walk(node.Left, target-node.Val) || walk(node.Right, target-node.Val) {
        return true
    }

    return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    return walk(root, targetSum)
}
