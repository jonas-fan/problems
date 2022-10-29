/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rrotate(node *TreeNode) *TreeNode {
    if node == nil || node.Left == nil {
        return node
    }

    head, target := node.Left, node.Left

    node.Left = nil

    for target.Right != nil {
        target = target.Right
    }

    target.Right = node

    return head
}

func increasingBST(root *TreeNode) *TreeNode {
    rotated := rrotate(root)

    if rotated == nil {
        return root
    } else if rotated != root {
        return increasingBST(rotated)
    }

    rotated.Right = increasingBST(rotated.Right)

    return rotated
}
