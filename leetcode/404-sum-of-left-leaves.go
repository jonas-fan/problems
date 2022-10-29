/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode, left bool) int {
    if node == nil {
        return 0
    }

    if node.Left == nil && node.Right == nil {
        if !left {
            return 0
        }

        return node.Val
    }

    return walk(node.Left, true) + walk(node.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
    return walk(root, false)
}
