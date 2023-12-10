/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
    return sum(root, low, high)
}

func sum(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }

    if root.Val < low {
        return sum(root.Right, low, high)
    } else if root.Val > high {
        return sum(root.Left, low, high)
    }

    return root.Val + sum(root.Left, low, high) + sum(root.Right, low, high)
}
