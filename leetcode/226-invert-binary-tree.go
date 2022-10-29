/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    left, right := node.Left, node.Right

    node.Left = invertTree(right)
    node.Right = invertTree(left)

    return node
}
