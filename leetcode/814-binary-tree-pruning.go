/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    if node.Left != nil || node.Right != nil {
        node.Left = walk(node.Left)
        node.Right = walk(node.Right)
    }

    if node.Left != nil || node.Right != nil {
        return node
    } else if node.Val != 0 {
        return node
    }

    return nil
}

func pruneTree(root *TreeNode) *TreeNode {
    return walk(root)
}
