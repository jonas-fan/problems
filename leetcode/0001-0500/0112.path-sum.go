// #tree #binary-tree

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

    target -= node.Val

    return walk(node.Left, target) ||
           walk(node.Right, target)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    return walk(root, targetSum)
}
